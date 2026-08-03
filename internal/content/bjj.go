package content

import (
	"net/url"
	"path"
)

type BJJTechnique struct {
	Name        string
	Position    string
	YouTube     string
	Description string
}

// Thumbnail is the preview image for a technique's video, or "" if it has none.
// hqdefault exists for every video including Shorts; maxresdefault often 404s.
func (t BJJTechnique) Thumbnail() string {
	id := videoID(t.YouTube)
	if id == "" {
		return ""
	}
	return "https://img.youtube.com/vi/" + id + "/hqdefault.jpg"
}

// videoID handles the two link shapes in the data: /watch?v=<id> and /shorts/<id>.
func videoID(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	if id := u.Query().Get("v"); id != "" {
		return id
	}
	if base := path.Base(u.Path); base != "." && base != "/" {
		return base
	}
	return ""
}

var BJJTechniques = []BJJTechnique{
	{
		Name:     "Kimura",
		Position: "north south",
		YouTube:  "https://www.youtube.com/shorts/UW62mtPhtW8",
	},
	{
		Name:     "Arm drag with stumble",
		Position: "standing",
		YouTube:  "https://www.youtube.com/watch?v=_uJeSFxjquM",
	},
	{
		Name:     "Trip from Uchi Mata",
		Position: "standing",
		YouTube:  "https://www.youtube.com/watch?v=YmLXLpcQpuQ",
	},
	{
		Name:     "Osoto to Sasae",
		Position: "standing",
		YouTube:  "https://www.youtube.com/shorts/oxPKRTVYVZQ",
	},
	{
		Name:     "Foot Sweep",
		Position: "standing",
		YouTube:  "https://www.youtube.com/shorts/wAPI3OnN3-Q",
	},
	{
		Name:        "Rear body lock foot sweep",
		Position:    "standing",
		Description: "Make hips perpendicular to partner's hips. Take two steps forward starting with the inside leg. Partner will have to adjust or you can just drag them down. While completing the second step turn into partner and have your inside foot catch your partner's outside foot and pull them down.",
	},
	{
		Name:        "Arm drag from bottom half",
		Position:    "bottom half guard",
		Description: "Put your partner's top hand in a modified kimura grip (e.g. if you're laying on your right side, grab their right wrist with your right hand, then your left hand goes behind their right arm to reinforce the grip). You threaten a Kimura and partner will posture to avoid. You follow up with them and straighten your arm to keep their arm trapped. The hand that was holding the wrist can switch to an arm drag.",
	},
	{
		Name:        "Passing half guard leg scoop",
		Position:    "top half guard",
		Description: "Perpendicular to your partner. Push your elbow behind their head to keep them crunched in a ball. Other arm attaches to the ankle of the leg that is making the knee shield. Use this leverage to clear the knee shield.",
	},
	{
		Name:     "Gift wrap choke",
		Position: "mount",
		YouTube:  "https://www.youtube.com/shorts/8HiXiADOwjE",
	},
	{
		Name:     "De-Ashi Barai",
		Position: "standing",
		YouTube:  "https://www.youtube.com/watch?v=Z3j4q0UWvFc",
	},
	{
		Name:     "Side Control D'Arce Choke",
		Position: "side control",
		YouTube:  "https://www.youtube.com/watch?v=GRNG1S2wIHE",
	},
	{
		Name:     "Side Control no Gi Baseball",
		Position: "side control",
		YouTube:  "https://www.youtube.com/watch?v=ChQKMVeofgk",
	},
}
