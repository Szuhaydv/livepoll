<script>
	export let url;
	export let text;
</script>

<a
	href={url}
	class="text-white visited:text-white hover:no-underline rounded-3xl border-white border-4 bg-gradient-to-b from-lime-300 to-green-700 w-max px-16 inline-block shadow-lg focus:border-white active:scale-105 transition"
>
	{text}
</a>
