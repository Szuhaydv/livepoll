<script>
	import Note from "../components/Note.svelte";
	import Button from "../components/Button.svelte";
</script>

<h1 class="text-8xl text-yellow-400 text-center pt-[4vh] font-actionJackson">
	Livepoll
</h1>
<Note title="Your poll is available at:">
	<div
		class="mt-14 flex items-center justify-between pl-8 pr-4 py-1 bg-white rounded-2xl"
	>
		<p class="font-caveatBrush text-5xl inline-block">
			https://www.something.com
		</p>
		<img class="w-14 h-14" src="/assets/link.svg" alt="Link icon" />
	</div>
	<div class="w-full flex justify-center mt-10">
		<Button text="Copy" url="https://www.something.com"></Button>
	</div>
</Note>
