package content

type Certificate struct {
	Name        string
	Date        string
	Image       string
	Description string
}

var Certificates = []Certificate{
	{
		Name:        "FastBit Embedded Brain Academy - Microcontroller Embedded C Programming",
		Date:        "Aug 2026",
		Image:       "/static/img/certificate/embedded_c.jpg",
		Description: "Embedded C programming for microcontrollers, from the ground up",
	},
	{
		Name:        "MIT - Data Science and Machine Learning",
		Date:        "Dec 2024",
		Image:       "/static/img/certificate/mit_data_science.jpg",
		Description: "Curriculum covers Data Science and ML skills needed to create business impact",
	},
	{
		Name:        "DataCamp - Data Scientist Professional Certificate",
		Date:        "Nov 2021",
		Image:       "/static/img/certificate/DataCamp_professional.jpg",
		Description: "Demonstrated ability to collect, analyze and interpret large amounts of data using machine learning",
	},
	{
		Name:        "DataCamp - Machine Learning",
		Date:        "Feb 2021",
		Image:       "/static/img/certificate/datacamp_machine_learning.jpg",
		Description: "Essential skills for machine learning",
	},
	{
		Name:        "DataCamp - Data Analyst",
		Date:        "Feb 2021",
		Image:       "/static/img/certificate/datacamp_data_analyst.jpg",
		Description: "Essential skills for data analyst",
	},
	{
		Name:        "DataCamp - Python Programmer Track",
		Date:        "Feb 2021",
		Image:       "/static/img/certificate/datacamp_programmer.jpg",
		Description: "Python programming fundamentals",
	},
	{
		Name:        "IBM - Machine Learning",
		Date:        "Dec 2020",
		Image:       "/static/img/certificate/ibm_machine_learning.jpg",
		Description: "Prepare for a career in machine learning",
	},
	{
		Name:        "DeepLearning.AI - Deep Learning Specialization",
		Date:        "Nov 2020",
		Image:       "/static/img/certificate/deeplearning.jpg",
		Description: "Master the fundamentals of deep neural networks",
	},
	{
		Name:        "Google Cloud - Machine Learning with TensorFlow on Google Cloud Platform",
		Date:        "Oct 2020",
		Image:       "/static/img/certificate/google_cloud.jpg",
		Description: "Training and deploying models on Google Cloud",
	},
	{
		Name:        "Johns Hopkins - Master of Arts",
		Date:        "May 2019",
		Image:       "/static/img/certificate/diploma_jhu-0.jpg",
		Description: "Focused on international development and econometrics, with a particular emphasis on the post-Soviet sphere",
	},
	{
		Name:        "Russian Ministry of Education",
		Date:        "May 2018",
		Image:       "/static/img/certificate/language_exam-2.jpg",
		Description: "Demonstrated upper-intermediate level of Russian language in a government exam",
	},
	{
		Name:        "UC Riverside - TEFL",
		Date:        "Aug 2012",
		Image:       "/static/img/certificate/tefl.jpg",
		Description: "Teaching English as a Foreign Language",
	},
	{
		Name:        "UC Riverside - Bachelor of Arts",
		Date:        "May 2012",
		Image:       "/static/img/certificate/ucr.jpg",
		Description: "Double major in Music and Comparative Literature",
	},
}
