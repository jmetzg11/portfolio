<script>
	import { currentActivities, background } from './data';
	import TopPart from './TopPart.svelte';
	import InfoPart from './InfoPart.svelte';
	import SampleProjects from './SampleProjects.svelte';
	import Skills from './Skills.svelte';
</script>

<section
	class="bg-gradient-to-b from-gray-100 to-gray-300 min-h-screen flex flex-col items-center justify-start p-8 space-y-12"
>
	<TopPart />
	<InfoPart data={currentActivities} />
	<InfoPart data={background} />
	<SampleProjects />
	<Skills />
</section>
