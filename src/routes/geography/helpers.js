export const normalizeData = (data, year) => {
	const endYear = 2025;
	let filteredData;
	if (year !== 'all') {
		filteredData = data.filter((item) => item.year === year);
	} else {
		filteredData = data;
	}
	let combinedData = {};
	filteredData.forEach((item) => {
		const location = item.location;
		if (location in combinedData) {
			combinedData[location]['text'] += `, ${item.text}`;
			combinedData[location]['value'] += item.value;
			combinedData[location]['color'] = colors[endYear - item.year];
			combinedData[location]['years'] += `, ${item.year}`;
		} else {
			combinedData[location] = {
				...item,
				color: colors[endYear - item.year],
				years: `${item.year}`
			};
		}
	});

	const resultArray = Object.values(combinedData);
	const maxValue = Math.max(...resultArray.map((item) => item.value));
	const minValue = Math.min(...resultArray.map((item) => item.value));
	const minRange = 4;
	const maxRange = 40;
	const weight = 3;

	return resultArray.map((item) => {
		const normalizedValue =
			((parseFloat(item.value) - minValue) / (maxValue - minValue + weight)) *
				(maxRange - minRange) +
			minRange;
		return {
			...item,
			normalizedValue
		};
	});
};

export const colors = [
	'#2E86AB', // Light steel blue
	'#2A7EA6', // Slightly darker blue
	'#2676A1', // Medium blue
	'#226E9C', //
	'#1E6697', //
	'#1A5E92', //
	'#16568D', //
	'#124E88', //
	'#0E4683', //
	'#0A3E7E', // Dark blue
	'#063679', //
	'#022E74', //
	'#01266F', //
	'#011E6A', //
	'#011665', //
	'#010E60', //
	'#01065B', // Very dark blue
	'#000356' // Almost black blue
];
