<script>
	const skills = [
		{
			name: 'Frontend',
			skills: ['React', 'Svelte', 'Tailwind CSS', 'Bootstrap']
		},
		{
			name: 'Backend',
			skills: ['Gin', 'Django', 'Temporal', 'FastAPI', 'Express']
		},
		{
			name: 'IoT/Embedded',
			skills: ['ESP-IDF', 'Raspberry Pi']
		},
		{
			name: 'DevOps',
			skills: ['Git', 'AWS', 'Docker', 'Fly.io', 'Microservices']
		},
		{
			name: 'Data & ML',
			skills: ['Machine Learning', 'Statistics', 'Pandas']
		},
		{
			name: 'Messaging & APIs',
			skills: ['MQTT', 'REST', 'WebSockets']
		},
		{
			name: 'Languages',
			skills: ['Python', 'Go', 'JavaScript', 'C', 'TypeScript']
		},
		{
			name: 'Databases',
			skills: ['SQLite', 'PostgreSQL', 'MySQL', 'MongoDB']
		},
		{
			name: 'Other',
			skills: ['Russian', 'Spanish', 'BJJ', 'Dogs']
		}
	];
</script>

<div class="w-full max-w-5xl">
	<div>
		<div class="text-center text-3xl md:text-4xl lg:text-5xl font-bold text-gray-800 mb-6">
			Skills
		</div>
		<div
			class="bg-white p-8 rounded-lg shadow-lg border border-gray-200 grid grid-cols-1 sm:grid-cols-2 gap-4 text-gray-700 text-lg"
		>
			{#each skills as skill}
				<div>
					<strong>{skill.name}:</strong>
					{skill.skills.join(', ')}
				</div>
			{/each}
		</div>
	</div>
</div>
