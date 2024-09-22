<script>
	export let name;
</script>

<svelte:head>
	<link
		href="https://fonts.cdnfonts.com/css/action-jackson"
		rel="stylesheet"
	/>
</svelte:head>

<main class="flex flex-col items-center">
	<h1
		class="text-8xl text-yellow-400 text-center mt-[20vh] font-actionJackson"
	>
		Welcome<br />to livepoll!
	</h1>
	<button
		class="text-white rounded-3xl border-4 border-white bg-gradient-to-b from-lime-300 to-green-700 w-max font-caveatBrush px-16 py-3 mt-40 text-[2.5rem] shadow-lg focus:border-white active:scale-105 transition"
	>
		New Poll
	</button>
</main>

<style>
	:global(body) {
		max-height: 100vh;
		margin: 0;
		padding: 0;
		background-image: url("/assets/tile.png");
		background-repeat: repeat;
		background-size: auto;
	}
</style>
