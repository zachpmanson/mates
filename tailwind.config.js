/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				'mate-red': '#d92b26',
				'mate-grey': '#212121',
				'mate-lime': '#c2d840'
			}
		}
	},
	plugins: []
};
