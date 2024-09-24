<script>
	import Note from "../components/Note.svelte";
	import Button from "../components/Button.svelte";
	let link = "https://www.something.com";

	function copyLink() {
		navigator.clipboard
			.writeText(link)
			.catch((err) => console.log("Colud not copy text", err));
	}
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note title="Your poll is available at:">
	<button
		class="flex items-center justify-between border-none w-full pl-8 pr-4 py-1 bg-white rounded-2xl cursor-pointer active:scale-95 transition"
		on:click={copyLink}
	>
		<span
			class="overflow-hidden whitespace-nowrap text-ellipsis inline-block"
			>{link}</span
		>
		<img class="w-14 h-14" src="/public/assets/link.svg" alt="Link icon" />
	</button>
	<div class="w-full flex justify-center mt-10">
		<Button text="Copy" actionButton="true" action={copyLink}></Button>
	</div>
</Note>
