<script>
	import Note from "../components/Note.svelte";

	import { onMount, onDestroy } from "svelte";
	import { fade } from "svelte/transition";

	class Poll {
		constructor(title, duration, options, createdAt) {
			this.title = title;
			this.duration = duration;
			this.options = options;
			this.createdAt = createdAt;
		}
	}

	let eventSource;
	export let params;

	onMount(() => {
		fetchPoll();
		const sseEndpoint = "http://localhost:7777/results/" + params.id,
			eventSource = new EventSource(sseEndpoint);

		eventSource.onmessage = (event) => {
			try {
				const data = JSON.parse(event.data);
				const id = poll.options.findIndex(
					(el) => el.id == data.option_id,
				);
				if (id == -1) {
					throw new Error("Update arrived for non-existing option");
				}
				poll.options[id].votes += 1;
				totalVotes += 1;
				calculatePercentages();
			} catch (error) {
				console.error("Error parsing JSON: ", error);
			}
		};

		eventSource.onerror = (error) => {
			console.error("EventSource failed:", error);
			eventSource.close();
		};
	});

	onDestroy(() => {
		if (eventSource) {
			eventSource.close();
		}
	});

	$: poll = new Poll("", 0, [], "");
	$: totalVotes = 0;
	let timeRemaining = null;
	$: inactivationDate = 0;
	let countdownRef;
	async function fetchPoll() {
		try {
			const getPollResponse = await fetch("/polls/" + params.id, {
				method: "GET",
			});
			if (!getPollResponse.ok) {
				const errMessage = await getPollResponse.text();
				console.error("Error getting poll: ", errMessage);
				return;
			}

			const data = await getPollResponse.json();
			poll = new Poll(
				data.title,
				data.duration,
				data.options,
				data.created_at,
			);
			for (const option of data.options) {
				totalVotes += option.votes;
			}
			calculatePercentages();
			initTimer();
		} catch (error) {
			console.error("Something went wrong: ", error);
		}
	}
	function initTimer() {
		calculateEndDate();
		countdownRef = setInterval(() => {
			updateTimeRemaining();
		}, 1000);
	}
	function calculatePercentages() {
		const isZero = totalVotes === 0;
		for (const option of poll.options) {
			if (isZero) {
				option.percentage = 0;
			} else {
				option.percentage = Math.round(
					(option.votes / totalVotes) * 100,
				);
			}
		}
	}
	function calculateEndDate() {
		const createdAt = new Date(poll.createdAt);
		const durationInParts = poll.duration.split(":").map(Number);
		const durationInMs =
			durationInParts[1] * 60000 + durationInParts[2] * 1000;
		inactivationDate = new Date(createdAt.getTime() + durationInMs);
	}
	function updateTimeRemaining() {
		const duration = Math.round(
			(inactivationDate.getTime() - Date.now()) / 1000,
		);
		if (duration < 0) {
			timeRemaining = 0;
			if (countdownRef) {
				clearInterval(countdownRef);
			}
			return;
		} else {
			timeRemaining = duration;
		}
	}

	onDestroy(() => {
		if (countdownRef) {
			clearInterval(countdownRef);
		}
	});
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note title={poll.title} titleMargin={2}>
	<div class="w-full flex justify-center items-center gap-5 mb-8">
		<img src="/public/assets/clock.svg" alt="Clock icon" />
		{#if timeRemaining !== null}
			<p transition:fade>
				{Math.floor(timeRemaining / 60)}:{String(
					timeRemaining % 60,
				).padStart(2, "0")}
			</p>
		{:else}
			<p class="opacity-0">0:00</p>
		{/if}
	</div>
	<ul class="flex flex-col">
		{#each poll.options as option, index}
			<li>
				<p
					class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12"
				>
					{option.name}
				</p>
				<div
					class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12 {index ==
					option.length - 1
						? 'border-b-2'
						: ''}"
				>
					<div
						class="loading-bar rounded-full h-8 mr-4 border border-black"
						style="width: {poll.options[index].percentage}%"
					></div>
					<span class="text-nowrap"
						>{poll.options[index].percentage} %</span
					>
				</div>
			</li>
		{/each}
	</ul>
</Note>

<style>
	.loading-bar {
		background-image: linear-gradient(
			-45deg,
			#ffc33c 0%,
			#ffc33c 25%,
			#ffaf2d 25%,
			#ffaf2d 50%,
			#ffc33c 50%,
			#ffc33c 75%,
			#ffaf2d 75%,
			#ffaf2d 100%
		);
		background-size: 2rem;
		animation: progress 0.75s linear infinite;
		box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
		transition: width 0.3s ease-in-out;
	}

	@keyframes progress {
		to {
			background-position-x: 2rem;
		}
	}
</style>
