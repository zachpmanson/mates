import { formSchema } from '$lib/components/report/schema';
import { error } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';

export const load = async () => {
	const form = await superValidate(
		{
			description: '',
			name: '',
			transport: 'tram'
		},
		zod(formSchema)
	);

	return { form };
};
