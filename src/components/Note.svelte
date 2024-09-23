<script>
	export let title;
	export let titleMargin = 10;
</script>

<div class="flex items-center justify-center">
	<main class="bg-yellow-paper w-[48rem] shadow-2xl py-12 px-24">
		<h2
			class="font-bangers text-5xl text-center"
			style="margin-bottom: {titleMargin / 4}rem"
		>
			{title}
		</h2>
		<slot></slot>
	</main>
</div>

<style>
	div {
		flex: 1;
	}
</style>
