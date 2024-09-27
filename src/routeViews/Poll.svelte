<script>
	import { onMount } from "svelte";
	import { location, push } from "svelte-spa-router";
	import { writable } from "svelte/store";
	import Button from "../components/Button.svelte";
	import Note from "../components/Note.svelte";

	export let params;

	onMount(async () => {
		if ($location != "/creator") {
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
				title.set(data.title);
				for (const option of data.options) {
					options.push(option.name);
					optionIDs.push(option.id);
				}
				options = options;
				optionIDs = optionIDs;
			} catch (error) {
				console.error("Something went wrong: ", error);
			}
		}
	});

	$: isCreatingPoll = $location === "/creator";
	const title = writable("");
	$: options = [];
	$: optionIDs = [];
	$: selectedOption = -1;
	function addOption() {
		options = [...options, ""];
	}
	function selectOption(index) {
		if (!isCreatingPoll) {
			selectedOption = index;
		}
	}
	function deleteOption(index) {
		options = [...options.slice(0, index), ...options.slice(index + 1)];
	}
	function focusOption(e) {
		e.focus();
	}
	function updateInput(e, index) {
		options[index] = e.target.value;
	}
	async function createPoll() {
		const formattedOptions = options.map((option) => ({ name: option }));
		try {
			const response = await fetch("/create-poll", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					title: $title,
					options: formattedOptions,
				}),
			});
			if (!response.ok) {
				const errMessage = await response.text();
				console.error("Error creating poll: ", errMessage);
			} else {
				const data = await response.json();
				push("/link/" + data.id);
			}
		} catch (error) {
			console.error("An error has happened: ", error);
		}
	}
	async function submitVote() {
		// handle when nothing is selected
		if (selectedOption == -1) {
			return;
		}
		try {
			const response = await fetch("/vote", {
				method: "POST",
				body: JSON.stringify({
					option_id: optionIDs[selectedOption],
					poll_id: params.id,
				}),
			});
			if (!response.ok) {
				const errMessage = await response.text();
				console.error("Error sending vote: ", errMessage);
			} else {
				push("/results/" + params.id);
			}
		} catch (error) {
			console.error("An error has happend: ", error);
		}
	}
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note isTitleEditable={isCreatingPoll} bind:title={$title}>
	<ul class="flex flex-col">
		{#each options as option, index}
			<li
				class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12 last:border-b-2"
			>
				<div
					class="relative h-10"
					on:click={() => selectOption(index)}
					on:keydown
				>
					<input
						type="radio"
						class="appearance-none w-10 h-10 rounded-full bg-white border border-black drop-shadow-lg"
						class:cursor-pointer={!isCreatingPoll}
						name="choice"
					/>
					{#if selectedOption == index}
						<div
							class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-4 h-4 rounded-full bg-black transition-transform duration-200 ease-in-out cursor-pointer"
							class:cursor-pointer={!isCreatingPoll}
						></div>
					{/if}
				</div>
				<input
					class="mx-6 text-ellipsis overflow-hidden whitespace-nowrap bg-transparent w-full border-none focus:outline-none disabled:text-black"
					type="text"
					value={option}
					use:focusOption={index}
					on:input={(e) => updateInput(e, index)}
					disabled={!isCreatingPoll}
				/>
				{#if isCreatingPoll}
					<button
						class="ml-auto mr-12 border-none"
						on:click={() => deleteOption(index)}
					>
						<img
							src="/public/assets/trash.svg"
							alt="Trash can icon"
						/>
					</button>
				{/if}
			</li>
		{/each}
	</ul>
	{#if isCreatingPoll}
		<div class="w-full flex justify-center mt-8">
			<button class="border-none" on:click={() => addOption()}>
				<img
					src="/public/assets/add.svg"
					alt="Add icon"
					class="w-10 h-10"
				/>
			</button>
		</div>
	{/if}
	<div class="w-full flex justify-center" slot="just-below">
		<Button
			actionButton={true}
			action={() => (isCreatingPoll ? createPoll() : submitVote())}
			text={isCreatingPoll ? "Create Poll" : "Submit"}
		></Button>
	</div>
</Note>
