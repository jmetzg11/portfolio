<script>
	const data = [
		{
			name: 'Kimura',
			position: 'north south',
			youtube: 'https://www.youtube.com/shorts/UW62mtPhtW8',
			description: null
		},
		{
			name: 'Arm drag with stumble',
			position: 'standing',
			youtube: 'https://www.youtube.com/watch?v=_uJeSFxjquM',
			description: null
		},
		{
			name: 'Trip from Uchi Mata',
			position: 'standing',
			youtube: 'https://www.youtube.com/watch?v=YmLXLpcQpuQ',
			description: null
		},
		{
			name: 'Foot Sweep',
			position: 'standing',
			youtube: 'https://www.youtube.com/shorts/wAPI3OnN3-Q',
			description: null
		},
		{
			name: 'Rear body lock foot sweep',
			position: 'standing',
			youtube: null,
			description:
				"Make hips perpendicular to partner's hips. Take to steps forward starting with the inside lget. Partner will have to adjust or you can just drag them down. While completing the second step turn into partner and have your inside foot catch your partner's outside foot and pull thme down."
		},
		{
			name: 'Arm drag from bottom half',
			position: 'bottom half guard',
			youtube: null,
			description:
				"Put your partner's top hand in a modified kimura grip (e.g. if you're laying on your right side, grab their right wrist with your write hand, then you left hand goes behind their right arm to reinforce the grip). You threaten a Kimura and partner will posture to avoide. You follow up with them and straighten your arm to keep their arm trapped. The hand that was holdding the wrist can switch to an arm drag."
		},
		{
			name: 'Passing half guard leg scoop',
			position: 'top half guard',
			youtube: null,
			description:
				'Perpendicular to you partners. Pust your elbow behind their head to keep them crunched in a ball. Other arm attacheds to the  ankle of the leg that is making the knee shield. You this leverage to clear the knee shield.'
		},
		{
			name: 'Gift wrap choke',
			position: 'mount',
			youtube: 'https://www.youtube.com/shorts/8HiXiADOwjE',
			description: null
		},
		{
			name: 'De-Ashi Barai',
			position: 'standing',
			youtube: 'https://www.youtube.com/watch?v=Z3j4q0UWvFc',
			description: null
		},
		{
			name: "Side Control D'Arce Choke",
			position: 'side control',
			youtube: 'https://www.youtube.com/watch?v=GRNG1S2wIHE',
			description: null
		},
		{
			name: 'Side Control no Gi Baseball',
			position: 'side control',
			youtube: 'https://www.youtube.com/watch?v=ChQKMVeofgk',
			description: null
		}
	];
</script>

<div
	class="p-5 bg-gradient-to-b min-h-screen min-w-screen from-gray-100 to-gray-300 grid grid-cols-2 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 gap-4"
>
	{#each data as item}
		<div class="bg-white rounded-lg shadow-md p-4 hover:shadow-lg transition-shadow">
			<h2 class="text-lg font-bold mb-2">{item.name}</h2>
			<p class="text-gray-600 mb-3">{item.position}</p>
			{#if item.youtube}
				<a
					href={item.youtube}
					target="_blank"
					rel="noopener noreferrer"
					class="inline-block px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors"
				>
					Youtube
				</a>
			{:else}
				<p class="text-gray-700">{item.description}</p>
			{/if}
		</div>
	{/each}
</div>
