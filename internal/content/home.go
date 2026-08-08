package content

const GithubIconPath = "M12 2C6.477 2 2 6.484 2 12.017c0 4.418 2.867 8.167 6.839 9.49.5.091.682-.217.682-.483 0-.237-.009-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.154-1.11-1.461-1.11-1.461-.908-.621.069-.608.069-.608 1.004.07 1.532 1.033 1.532 1.033.892 1.529 2.341 1.087 2.91.831.091-.647.349-1.087.635-1.337-2.221-.253-4.555-1.114-4.555-4.951 0-1.093.39-1.987 1.029-2.686-.103-.253-.446-1.271.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.338 1.908-1.296 2.748-1.026 2.748-1.026.546 1.379.203 2.397.1 2.65.64.7 1.028 1.593 1.028 2.686 0 3.847-2.337 4.695-4.566 4.944.358.309.678.919.678 1.852 0 1.336-.012 2.415-.012 2.743 0 .268.18.579.688.481A10.019 10.019 0 0022 12.017C22 6.484 17.523 2 12 2z"

type Project struct {
	Name        string
	URL         string
	Description string
	CodeURL     string
	External    bool
}

var Projects = []Project{
	{
		Name:        "See My Family",
		URL:         "https://seemyfamily.net",
		Description: "A site for my family to track and visualize relations across the globe and time",
		CodeURL:     "https://github.com/jmetzg11/seemyfamily_backend",
		External:    true,
	},
	{
		Name:        "Automated investing",
		URL:         "/investing",
		Description: "A passive way to invest",
	},
	{
		Name:        "Farmsville",
		URL:         "https://farmsville.fly.dev/",
		Description: "A social network for farmers",
		CodeURL:     "https://github.com/jmetzg11/farmsville",
		External:    true,
	}, 
	{
		Name:        "Geography",
		URL:         "/geography",
		Description: "A tracker to see all the places I have visited",
	},
	{
		Name:        "Government Spending",
		URL:         "https://gov-spending.fly.dev/",
		Description: "Visualize government spending by agency and geography",
		CodeURL:     "https://github.com/jmetzg11/gov-spending",
		External:    true,
	},
	{
		Name:        "Helen's Books",
		URL:         "https://helenmetzgerbook.com/",
		Description: "under construction",
		CodeURL:     "https://github.com/jmetzg11/helen_portfolio",
		External:    true,
	},
	{
		Name:        "BJJ Reference",
		URL:         "/bjj",
		Description: "A reference for my favorite techniques so I don't forget them",
	},
	{
		Name:        "Education & Certificates",
		URL:         "/certificates",
		Description: "My university diplomas and related certificates",
	},
}
