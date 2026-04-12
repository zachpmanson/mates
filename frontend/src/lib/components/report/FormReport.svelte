<script lang="ts">
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';
	import { formSchema } from './schema';
	import Input from '../ui/input/input.svelte';
	import * as Form from '../ui/form';

	export let onSuccess: () => void = () => {};

	let locationStatus: 'idle' | 'loading' | 'ok' | 'error' = 'idle';

	const sf = superForm(
		defaults({ title: '', summary: '', lat: 0, long: 0 }, zod(formSchema)),
		{
			SPA: true,
			validators: zod(formSchema),
			onUpdate: async ({ form }) => {
				if (!form.valid) return;
				onSuccess();
			}
		}
	);

	const { form, enhance, submitting } = sf;

	function locateMe() {
		locationStatus = 'loading';
		navigator.geolocation.getCurrentPosition(
			(pos) => {
				$form.lat = pos.coords.latitude;
				$form.long = pos.coords.longitude;
				locationStatus = 'ok';
			},
			() => {
				locationStatus = 'error';
			}
		);
	}
</script>

<form method="POST" use:enhance class="flex flex-col gap-4">
	<Form.Field form={sf} name="summary">
		<Form.Control let:attrs>
			<Form.Label>Summary</Form.Label>
			<Input {...attrs} bind:value={$form.summary} placeholder="Describe the sighting…" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field form={sf} name="title">
		<Form.Control let:attrs>
			<Form.Label>Title <span class="text-gray-400 font-normal">(optional)</span></Form.Label>
			<Input {...attrs} bind:value={$form.title} placeholder="Short title" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field form={sf} name="lat">
		<Form.Control let:attrs>
			<Form.Label>Location</Form.Label>
			<div class="flex gap-2 items-center">
				<Input {...attrs} type="number" step="any" bind:value={$form.lat} placeholder="Lat" />
				<Input
					name="long"
					type="number"
					step="any"
					bind:value={$form.long}
					placeholder="Long"
				/>
				<button
					type="button"
					on:click={locateMe}
					class="shrink-0 border rounded px-3 py-1.5 text-sm hover:bg-gray-50"
				>
					{locationStatus === 'loading' ? 'Locating…' : 'Use my location'}
				</button>
			</div>
			{#if locationStatus === 'error'}
				<p class="text-sm text-red-500">Could not get location.</p>
			{/if}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Button disabled={$submitting}>
		{$submitting ? 'Submitting…' : 'Report sighting'}
	</Form.Button>
</form>
