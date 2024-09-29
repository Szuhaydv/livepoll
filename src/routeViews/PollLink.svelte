<script>
	import Note from "../components/Note.svelte";
	import Button from "../components/Button.svelte";
	import { push } from "svelte-spa-router";
	export let params;

	const link = "http://localhost:7777/#/fill-poll/" + params.id;

	function copyLink() {
		navigator.clipboard
			.writeText(link)
			.catch((err) => console.log("Could not copy text", err));
	}

	function navigate() {
		push("/fill-poll/" + params.id);
	}
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note title="Your poll is available at:">
	<button
		class="flex items-center justify-between border-none w-full pl-8 pr-4 py-1 bg-white rounded-2xl cursor-pointer active:scale-95 transition"
		on:click={navigate}
	>
		<span
			class="overflow-hidden whitespace-nowrap text-ellipsis inline-block"
			>{link}</span
		>
		<img class="w-10 h-10" src="/public/assets/link.svg" alt="Link icon" />
	</button>
	<div class="w-full flex justify-center mt-10">
		<Button text="Copy" actionButton="true" action={copyLink}></Button>
	</div>
</Note>
