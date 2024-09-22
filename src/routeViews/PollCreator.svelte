<script>
	const texts = ["Situps", "Push ups", "Squats", "Pull ups"];
</script>

<h1 class="text-8xl text-yellow-400 text-center mt-[4vh] font-actionJackson">
	Livepoll
</h1>
<main
	class="bg-yellow-paper w-[54rem] h-[30rem] shadow-2xl mx-auto mt-[16vh] py-12 px-32"
>
	<h2 class="font-bangers text-5xl text-center">
		What exercise should I do?
	</h2>
	<ul class="flex flex-col gap-8 mt-16">
		{#each texts as text}
			<li
				class="before:content-[''] before:w-10 before:h-10 before:rounded-full before:bg-white before:absolute before:top-1/2 before:right-[102%] relative before:translate-y-[-50%] before:border-2 before:border-black font-caveatBrush text-3xl"
			>
				{text}
			</li>
		{/each}
	</ul>
</main>
