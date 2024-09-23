<script>
	import Button from "../components/Button.svelte";
	const texts = ["Situps", "Push ups", "Squats", "Pull ups"];
	$: selectedOption = 0;
</script>

<h1 class="text-8xl text-yellow-400 text-center mt-[4vh] font-actionJackson">
	Livepoll
</h1>
<main
	class="bg-yellow-paper w-[54rem] h-[30rem] shadow-2xl mx-auto mt-[16vh] py-12 px-32 mb-10"
>
	<h2 class="font-bangers text-5xl text-center">
		What exercise should I do?
	</h2>
	<ul class="flex flex-col mt-16">
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
						class="appearance-none w-10 h-10 rounded-full bg-white border-2 border-black cursor-pointer"
						name="choice"
					/>
					{#if selectedOption == index}
						<div
							class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-4 h-4 rounded-full bg-black transition-transform duration-200 ease-in-out cursor-pointer"
						></div>
					{/if}
				</div>
				<p class="font-caveatBrush text-3xl ml-6">
					{text}
				</p>
			</li>
		{/each}
	</ul>
</main>
<div class="w-full flex justify-center">
	<Button url="#/link" text="Create Poll"></Button>
</div>
