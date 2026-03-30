<script lang="ts">
	import logo from '$lib/images/logo.svg';

	import { formSchema, type FormSchema } from '$lib/components/report/schema';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Drawer from '$lib/components/ui/drawer';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Select from '$lib/components/ui/select';
	import { Textarea } from '$lib/components/ui/textarea';
	import { zod } from 'sveltekit-superforms/adapters';
	import { superForm } from 'sveltekit-superforms/client';

	const { form, errors, message, constraints, enhance } = superForm(
		{
			description: '',
			name: '',
			transport: 'tram'
		},
		{
			SPA: true,
			validators: zod(formSchema),
			onUpdate({ form }) {
				if (form.valid) {
					// TODO: Call an external API with form.data, await the result and update form
					console.log('Valid');
				}
			}
		}
	);

	function handleSubmit() {
		console.log('submit');
		// open = false;
	}

	let open = false;
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
				<Drawer.Description></Drawer.Description>
			</Drawer.Header>
			<div class="p-2">
				<form method="POST" use:enhance>
					<div class="flex gap-2 flex-col">
						<Label for="message">Where did you see this mate?</Label>
						<Textarea
							placeholder="Mate spotted at..."
							id="message"
							bind:value={$form.description}
						/>

						<Select.Root
							selected={{
								value: $form.transport,
								label: $form.transport.charAt(0).toUpperCase() + $form.transport.slice(1)
							}}
							onSelectedChange={(v) => {
								v && ($form.transport = v.value);
							}}
						>
							<Select.Trigger>
								<Select.Value placeholder="" />
							</Select.Trigger>
							<Select.Content>
								<Select.Item value="tram" label="Tram" />
								<Select.Item value="train" label="Train" />
								<Select.Item value="bus" label="Bus" />
							</Select.Content>
						</Select.Root>

						<Label for="name">Name</Label>
						<Input id="name" placeholder="Your name (optional)" bind:value={$form.name} />
					</div>
				</form>
			</div>

			<Drawer.Footer>
				<Drawer.Close class="w-full">
					<Button
						class="w-full"
						variant={'action'}
						disabled={!!$errors.description}
						on:click={handleSubmit}>Submit</Button
					>
				</Drawer.Close>

				<Drawer.Close>Cancel</Drawer.Close>
			</Drawer.Footer>
		</Drawer.Content>
	</Drawer.Root>
</header>
