<script>
	export let title;
</script>

<main
	class="bg-yellow-paper w-[54rem] shadow-2xl mx-auto mt-[16vh] py-12 px-32 mb-10"
>
	<h2 class="font-bangers text-5xl text-center">
		{title}
	</h2>
	<slot></slot>
</main>
