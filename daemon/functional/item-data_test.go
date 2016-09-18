package functional

var testItems itemList = itemList{
	{
		InternalName: "Item1",
		InternalDesc: "Item1Desc",
		Type:         "ITEM",
		Name:         "Item 1",
		Desc:         "Item 1 Desc",
		MaxCount:     100,
		Metadata:     map[string]string{"foo": "bar", "feh": "baz"},
		Flags:        0,
		Language:     "en",
	},
	{
		InternalName: "Item2",
		InternalDesc: "Item2Desc",
		Type:         "ITEM",
		Name:         "Item 2",
		Desc:         "Item 2 Desc",
		MaxCount:     100,
		Metadata:     nil,
		Flags:        0,
		Language:     "en",
	},
	{
		InternalName: "Item3",
		InternalDesc: "Item3Desc",
		Type:         "ITEM",
		Name:         "Item 3",
		Desc:         "Item 3 Desc",
		MaxCount:     100,
		Metadata:     nil,
		Flags:        0,
		Language:     "en",
	},
	{
		InternalName: "GlobalItem1",
		InternalDesc: "GlobalItem1Desc",
		Type:         "ITEM",
		Name:         "Global Item 1",
		Desc:         "Global Item 1 Desc",
		MaxCount:     100,
		Metadata:     map[string]string{"foo": "bar", "feh": "baz"},
		Flags:        1,
		Language:     "en",
	},
	{
		InternalName: "Currency1",
		InternalDesc: "Currency1Desc",
		Type:         "CURRENCY",
		Name:         "Currency 1",
		Desc:         "Currency 1 Desc",
		MaxCount:     9999999999999,
		Flags:        0,
		Language:     "en",
	},
	{
		InternalName: "XP",
		InternalDesc: "XPDesc",
		Type:         "CURRENCY",
		Name:         "XP Name",
		Desc:         "XP Desc",
		MaxCount:     9999999999999,
		Flags:        1,
		Language:     "en",
	},
	{
		InternalName: "Skill1",
		InternalDesc: "Skill1Desc",
		Type:         "SKILL",
		Name:         "Skill 1",
		Desc:         "Skill 1 Desc",
		MaxCount:     1000,
		Flags:        0,
		Language:     "en",
	},
	{
		InternalName: "Skill2",
		InternalDesc: "Skill2Desc",
		Type:         "SKILL",
		Name:         "Skill 2",
		Desc:         "Skill 2 Desc",
		MaxCount:     1000,
		Flags:        0,
		Language:     "en",
	},
	{
		InternalName: "GlobalSkill1",
		InternalDesc: "GlobalSkill1Desc",
		Type:         "SKILL",
		Name:         "GobalS kill 1",
		Desc:         "Global Skill 1 Desc",
		MaxCount:     1000,
		Flags:        1,
		Language:     "en",
	},
}
