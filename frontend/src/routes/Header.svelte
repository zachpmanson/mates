<script lang="ts">
	import logo from '$lib/images/logo.svg';

	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import { formSchema } from '$lib/components/report/schema';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Drawer from '$lib/components/ui/drawer';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Textarea } from '$lib/components/ui/textarea';
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';

	let locationStatus: 'idle' | 'loading' | 'ok' | 'error' = 'idle';
	let open = false;

	const sf = superForm(
		defaults({ title: '', summary: '', lat: 0, long: 0 }, zod(formSchema)),
		{
			SPA: true,
			validators: zod(formSchema),
			onUpdate: async ({ form }) => {
				if (!form.valid) return;
				const feedId = $page.url.searchParams.get('feed_id');
				if (!feedId) return;
				const res = await fetch('/api/sightings', {
					method: 'POST',
					headers: { 'Content-Type': 'application/json' },
					body: JSON.stringify({
						title: form.data.title || undefined,
						summary: form.data.summary,
						lat: form.data.lat,
						long: form.data.long,
						feed_id: parseInt(feedId)
					})
				});
				if (res.ok) {
					open = false;
					reset();
					locationStatus = 'idle';
					invalidateAll();
				}
			}
		}
	);

	const { form, errors, enhance, submitting, reset } = sf;

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

<header class="flex justify-between items-center gap-1 p-3 bg-mate-grey">
	<div class="h-8">
		<a
			href="/"
			class="flex gap-2 h-full items-center uppercase font-light text-white hover:text-mate-lime transition duration-75"
		>
			<img src={logo} alt="Where are our mates?" class="h-full" />
			Where Are Our Mates?
		</a>
	</div>

	<Drawer.Root bind:open>
		<Drawer.Trigger>
			<Button variant={'action'}>Report Mate Sighting</Button>
		</Drawer.Trigger>
		<Drawer.Content class="max-w-lg m-auto">
			<Drawer.Header>
				<Drawer.Title>Mate Spotted!</Drawer.Title>
			</Drawer.Header>
			<div class="p-4">
				<form method="POST" use:enhance class="flex flex-col gap-3">
					<div>
						<Label for="summary">Where did you see this mate?</Label>
						<Textarea
							placeholder="Mate spotted at…"
							id="summary"
							bind:value={$form.summary}
							class="mt-1"
						/>
						{#if $errors.summary}
							<p class="text-sm text-red-500 mt-1">{$errors.summary}</p>
						{/if}
					</div>

					<div>
						<Label for="title">Title <span class="text-gray-400 font-normal">(optional)</span></Label>
						<Input id="title" placeholder="Short title" bind:value={$form.title} class="mt-1" />
					</div>

					<div>
						<Label>Location</Label>
						<div class="mt-1 flex gap-2 items-center">
							<Input type="number" step="any" bind:value={$form.lat} placeholder="Lat" />
							<Input type="number" step="any" bind:value={$form.long} placeholder="Long" />
							<button
								type="button"
								on:click={locateMe}
								class="shrink-0 border rounded px-3 py-1.5 text-sm hover:bg-gray-50"
							>
								{locationStatus === 'loading' ? 'Locating…' : 'Use my location'}
							</button>
						</div>
						{#if locationStatus === 'error'}
							<p class="text-sm text-red-500 mt-1">Could not get location.</p>
						{/if}
					</div>

					<Drawer.Footer class="px-0">
						<Button type="submit" class="w-full" variant={'action'} disabled={$submitting}>
							{$submitting ? 'Submitting…' : 'Submit'}
						</Button>
						<Drawer.Close>Cancel</Drawer.Close>
					</Drawer.Footer>
				</form>
			</div>
		</Drawer.Content>
	</Drawer.Root>
</header>
