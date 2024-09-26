<script>
	import Router from "svelte-spa-router";
	import Home from "./routeViews/Home.svelte";
	import Poll from "./routeViews/Poll.svelte";
	import PollResults from "./routeViews/PollResults.svelte";
	import PollLink from "./routeViews/PollLink.svelte";

	const routes = {
		"/": Home,
		"/creator": Poll,
		"/link/:id": PollLink,
		"/fill-poll/:id": Poll,
		"/results/:id": PollResults,
	};
</script>

<Router {routes} />

<style>
	:global(body) {
		max-height: 100vh;
		margin: 0;
		padding: 0;
		background-image: url("/public/assets/tile.png");
		background-repeat: repeat;
		background-size: auto;
		font-family: "Caveat Brush";
		font-size: 2rem;
		display: flex;
		flex-direction: column;
	}
</style>
