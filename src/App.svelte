<script>
	import Router from "svelte-spa-router";
	import Home from "./routeViews/Home.svelte";
	import PollCreator from "./routeViews/PollCreator.svelte";
	import PollResults from "./routeViews/PollResults.svelte";
	import PollLink from "./routeViews/PollLink.svelte";

	const routes = {
		"/": Home,
		"/creator": PollCreator,
		"/results": PollResults,
		"/link": PollLink,
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
