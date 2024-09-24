<script>
	import Button from "../components/Button.svelte";
	import Note from "../components/Note.svelte";
	$: texts = ["Situps", "Push ups", "Squats", "Pull ups"];
	$: selectedOption = -1;
	function addOption() {
		texts = [...texts, ""];
	}
	function deleteOption(index) {
		texts = [...texts.slice(0, index), ...texts.slice(index + 1)];
	}
	function focusOption(e) {
		e.focus();
	}
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note title="What exercise should I do?">
	<ul class="flex flex-col">
		{#each texts as text, index}
			<li
				class="flex items-center border-t-2 h-16 border-slate-400 w-full pl-12 last:border-b-2"
			>
				<div
					class="relative h-10"
					on:click={() => (selectedOption = index)}
					on:keydown
				>
					<input
						type="radio"
						class="appearance-none w-10 h-10 rounded-full bg-white border border-black cursor-pointer drop-shadow-lg"
						name="choice"
					/>
					{#if selectedOption == index}
						<div
							class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-4 h-4 rounded-full bg-black transition-transform duration-200 ease-in-out cursor-pointer"
						></div>
					{/if}
				</div>
				<input
					class="mx-6 text-ellipsis overflow-hidden whitespace-nowrap bg-transparent w-full border-none focus:outline-none"
					type="text"
					value={text}
					use:focusOption={index}
				/>
				<button
					class="ml-auto mr-12 border-none"
					on:click={() => deleteOption(index)}
				>
					<img src="/public/assets/trash.svg" alt="Trash can icon" />
				</button>
			</li>
		{/each}
	</ul>
	<div class="w-full flex justify-center mt-8">
		<button class="border-none" on:click={() => addOption()}>
			<img
				src="/public/assets/add.svg"
				alt="Add icon"
				class="w-10 h-10"
			/>
		</button>
	</div>
	<div class="w-full flex justify-center" slot="just-below">
		<Button url="#/link" text="Create Poll"></Button>
	</div>
</Note>
