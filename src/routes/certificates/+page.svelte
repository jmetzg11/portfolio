<script>
	const data = [
		{
			name: 'MIT - Data Science and Machine Learning',
			date: 'Dec 2024',
			img: '/certificate/mit_data_science.jpg',
			description: 'Curriculum convers Data Science and ML skills needed to create business impact'
		},
		{
			name: 'DataCamp - Data Scientist Professional Certificate',
			date: 'Nov 2021',
			img: '/certificate/DataCamp_professional.jpg',
			description:
				'Demostrated ability to collect, analyze and interpert large amounts of data using machine learning'
		},
		{
			name: 'DataCamp - Machine Learning',
			date: 'Feb 2021',
			img: 'certificate/datacamp_machine_learning.jpg',
			description: 'Essential skills for machine learning'
		},
		{
			name: 'DataCamp - Data Analyst',
			date: 'Feb 2021',
			img: '/certificate/datacamp_data_analyst.jpg',
			description: 'Essential skills for data analyst'
		},
		{
			name: 'DataCamp - Python Programmer Track',
			date: 'Feb 2021',
			img: 'certificate/datacamp_programmer.jpg',
			description: 'Python programming fundamentals'
		},
		{
			name: 'IBM - Machine Learning',
			date: 'Dec 2020',
			img: '/certificate/ibm_machine_learning.jpg',
			description: 'Prepare for a career in machine learning'
		},
		{
			name: 'DeepLearning.AI - Deep Learning Specialization',
			date: 'Nov 2020',
			img: 'certificate/deeplearning.jpg',
			description: 'Master the fundamentals of deep neural networks'
		},
		{
			name: 'Google Cloud - Machine Learning with TensorFlow on Google Cloud Platform',
			date: 'Oct 2020',
			img: '/certificate/google_cloud.jpg',
			description: 'Training and deploying models on Google Cloud'
		},
		{
			name: 'Johns Hopkins - Master of Arts',
			date: 'May 2019',
			img: '/certificate/diploma_jhu-0.jpg',
			description:
				'Focused on international development and econometrics, with a particular emphasis on the post-Soviet sphere'
		},
		{
			name: 'Russian Ministry of Education',
			date: 'May 2018',
			img: '/certificate/language_exam-2.jpg',
			description: 'Demonstrated upper-intermediate level of Russian language in a government exam'
		},
		{
			name: 'UC Riverside - TEFL',
			date: 'Aug 2012',
			img: '/certificate/tefl.jpg',
			description: 'Teaching English as a Foreign Language'
		},
		{
			name: 'UC Riverside - Bachelor of Arts',
			date: 'May 2012',
			img: '/certificate/ucr.jpg',
			description: 'Double major in Music and Comparative Literature'
		}
	];
</script>

<div class="min-h-screen bg-gray-50 py-12">
	<div class="container mx-auto max-w-4xl px-4">
		<div class="space-y-8">
			{#each data as d, i}
				<div class="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-200">
					<div class="p-8">
						<div class="flex justify-between items-start mb-6">
							<h3 class="text-2xl font-semibold text-gray-900">{d.name}</h3>
							<span class="text-sm font-medium text-gray-500 bg-gray-100 px-3 py-1 rounded-full">
								{new Date(d.date).toLocaleDateString('en-US', {
									month: 'short',
									year: 'numeric'
								})}
							</span>
						</div>
						<p class="text-gray-600 mb-6 text-lg">{d.description}</p>
						<div class="relative group">
							<img
								class="w-full h-auto rounded-lg shadow-sm"
								src={d.img}
								alt={`Certificate: ${d.name}`}
							/>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
