<script>
	const projects = [
		{
			name: 'See My Family',
			link: 'https://seemyfamily.net',
			description:
				'A site for my family to track and visualize relations across the globe and time',
			code: 'https://github.com/jmetzg11/seemyfamily_backend'
		},
		{
			name: 'ETF Investing',
			link: 'https://stocks-etf.fly.dev/',
			description:
				'An automated investment strategy that targets the globally 11 sectors of the economy',
			code: 'https://github.com/jmetzg11/stocks-etf'
		},
		{
			name: 'Farmsville',
			link: 'https://farmsville.fly.dev/',
			description: 'A social network for farmers',
			code: 'https://github.com/jmetzg11/farmsville'
		},
		{
			name: 'Government Spending',
			link: 'https://gov-spending.fly.dev/',
			description: 'Visualize government spending by agency and geography',
			code: 'https://github.com/jmetzg11/gov-spending'
		},
		{
			name: "Helen's Books",
			link: 'https://helenmetzgerbook.com/',
			description: 'under constructions',
			code: 'https://github.com/jmetzg11/helen_portfolio'
		},
		{
			name: 'BJJ Reference',
			link: 'bjj',
			description: "A reference for my favorite techniques so I don't forget them",
			code: false
		},
		{
			name: 'Geography',
			link: 'geography',
			description: 'A tracker to see all the places I have visited',
			code: false
		},

		{
			name: 'Education & Certificates',
			link: 'certificates',
			description: 'My university diplomas and related certificates',
			code: false
		}
	];

	const github =
		'M12 2C6.477 2 2 6.484 2 12.017c0 4.418 2.867 8.167 6.839 9.49.5.091.682-.217.682-.483 0-.237-.009-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.154-1.11-1.461-1.11-1.461-.908-.621.069-.608.069-.608 1.004.07 1.532 1.033 1.532 1.033.892 1.529 2.341 1.087 2.91.831.091-.647.349-1.087.635-1.337-2.221-.253-4.555-1.114-4.555-4.951 0-1.093.39-1.987 1.029-2.686-.103-.253-.446-1.271.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.338 1.908-1.296 2.748-1.026 2.748-1.026.546 1.379.203 2.397.1 2.65.64.7 1.028 1.593 1.028 2.686 0 3.847-2.337 4.695-4.566 4.944.358.309.678.919.678 1.852 0 1.336-.012 2.415-.012 2.743 0 .268.18.579.688.481A10.019 10.019 0 0022 12.017C22 6.484 17.523 2 12 2z';

	const isInternalLink = (link) => {
		return !link.startsWith('http');
	};
</script>

<div class="w-full max-w-5xl">
	<div class="text-center text-3xl md:text-4xl lg:text-5xl font-bold text-gray-800 mb-6">
		Sample Projects
	</div>
	<ul class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
		{#each projects as project}
			<li
				class="p-6 bg-white rounded-lg shadow-lg border border-gray-200 hover:shadow-2xl transition-shadow transform hover:-translate-y-1 flex flex-col"
			>
				<a
					href={project.link}
					target="_blank"
					rel="noopener noreferrer"
					class="text-xl font-semibold text-blue-600 hover:text-blue-800 hover:underline"
				>
					{project.name}
				</a>
				<p
					class="mt-2 text-gray-600 text-sm
				flex-grow"
				>
					{project.description}
				</p>
				{#if project.code}
					<a
						href={project.code}
						target="_blank"
						rel="noopener noreferrer"
						class="inline-flex items-center gap-2 mt-4 text-gray-700 hover:text-black transition-all group px-2 py-1 rounded hover:bg-gray-100"
					>
						<span class="text-sm font-medium flex items-center gap-2">
							View code <svg
								xmlns="http://www.w3.org/2000/svg"
								class="h-6 w-6 transition-transform group-hover:scale-110"
								fill="currentColor"
								viewBox="0 0 24 24"
							>
								<path d={github} />
							</svg>
						</span>
					</a>
				{/if}
			</li>
		{/each}
	</ul>
</div>
