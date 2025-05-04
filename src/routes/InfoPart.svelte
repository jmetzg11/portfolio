<script>
	let { data } = $props();
</script>

<div class="w-full max-w-5xl">
	<div class="text-center text-3xl md:text-4xl lg:text-5xl font-bold text-gray-800 mb-6">
		{data.title}
	</div>
	<div class="bg-white p-8 rounded-lg shadow-lg border border-gray-200">
		<p class="text-lg text-gray-700 leading-relaxed">
			{data.text}
		</p>
	</div>
</div>
