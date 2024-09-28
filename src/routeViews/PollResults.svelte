<script>
	import Note from "../components/Note.svelte";

	import { onMount, onDestroy } from "svelte";
	import { writable } from "svelte/store";
	import { fade } from "svelte/transition";
	import { Poll, fetchPoll } from "../dataService.js";

	class Timer {
		constructor(timeRemaining, inactivationDate, countdownRef) {
			this.timeRemaining = timeRemaining;
			this.inactivationDate = inactivationDate;
			this.countdownRef = countdownRef;
		}
	}

	let eventSource;
	export let params;
	const poll = writable(new Poll());
	const timer = writable(new Timer());
	const totalVotes = writable(0);
	onMount(() => {
		fetchPoll(params.id, poll, timer, totalVotes);
		const sseEndpoint = "http://localhost:7777/results/" + params.id;
		eventSource = new EventSource(sseEndpoint);

		eventSource.onmessage = (event) => {
			try {
				const data = JSON.parse(event.data);
				const id = $poll.options.findIndex(
					(el) => el.id == data.option_id,
				);
				if (id == -1) {
					throw new Error("Update arrived for non-existing option");
				}
				$poll.options[id].votes += 1;
				totalVotes.update((n) => n + 1);
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
		if ($timer.countdownRef) {
			clearInterval($timer.countdownRef);
		}
		if (eventSource) {
			eventSource.close();
		}
	});
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note title={poll.title} titleMargin={2}>
	<div class="w-full flex justify-center items-center gap-5 mb-8">
		<img src="/public/assets/clock.svg" alt="Clock icon" />
		{#if $timer.timeRemaining !== undefined}
			<p transition:fade>
				{Math.floor($timer.timeRemaining / 60)}:{String(
					$timer.timeRemaining % 60,
				).padStart(2, "0")}
			</p>
		{:else}
			<p class="opacity-0">0:00</p>
		{/if}
	</div>
	<ul class="flex flex-col">
		{#each $poll.options as option, index}
			<li>
				<p
					class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12"
				>
					{option.name}
				</p>
				<div
					class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12 {index ==
					$poll.options.length - 1
						? 'border-b-2'
						: ''}"
				>
					<div
						class="loading-bar rounded-full h-8 mr-4 border border-black"
						style="width: {option.percentage}%"
					></div>
					<span class="text-nowrap">{option.percentage} %</span>
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
