export const requestMethods = [
	{
		value: "GET",
		title: "GET",
		primary_color: "oklch(79.2% 0.209 151.711)",
		on_primary_color: "green",
	},
	{
		value: "POST",
		title: "POST",
		primary_color: "oklch(82.8% 0.189 84.429)",
		on_primary_color: "#ffffff",
	},
	{
		value: "PUT",
		title: "PUT",
		primary_color: "oklch(74.6% 0.16 232.661)",
		on_primary_color: "#ffffff",
	},
	{
		value: "PATCH",
		title: "PATCH",
		primary_color: "oklch(71.4% 0.203 305.504)",
		on_primary_color: "#ffffff",
	},
	{
		value: "DELETE",
		title: "DELETE",
		primary_color: "oklch(70.4% 0.191 22.216)",
		on_primary_color: "#ffffff",
	},
	{
		value: "HEAD",
		title: "HEAD",
		primary_color: "oklch(71.8% 0.202 349.761)",
		on_primary_color: "#ffffff",
	},
	{
		value: "TRACE",
		title: "TRACE",
		primary_color: "oklch(77.7% 0.152 181.912)",
		on_primary_color: "#ffffff",
	},
	{
		value: "OPTIONS",
		title: "OPTIONS",
		primary_color: "oklch(76.5% 0.177 163.223)",
		on_primary_color: "#ffffff",
	},
];

export const getPrimaryColor = (list: any[], value: string) => {
	return list.find((m) => m.value === value)?.primary_color;
};

export const getOnPrimaryColor = (list: any[], value: string) => {
	return list.find((m) => m.value === value)?.on_primary_color;
};

export const getTitle = (list: any[], value: string) => {
	return list.find((m) => m.value === value)?.title;
};
