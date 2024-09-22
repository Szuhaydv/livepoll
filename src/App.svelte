<script>
	import Router from "svelte-spa-router";
	import Home from "./routeViews/Home.svelte";

	const routes = {
		"/": Home,
	};
</script>

<Router {routes} />

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
