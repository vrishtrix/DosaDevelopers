export const format = (seconds: number | string) => {
	if (isNaN(seconds as number)) return '0:00';

	const minutes = Math.floor((seconds as number) / 60);
	seconds = Math.floor((seconds as number) % 60);
	if (seconds < 10) seconds = `0${seconds}`;

	return `${minutes}:${seconds}`;
};
