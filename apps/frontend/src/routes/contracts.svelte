<script lang="ts">
	import PlusWhiteIcon from '../assets/icons/plus-white.svg';
	import ContractCard from '../lib/components/ContractCard.svelte';
</script>

<div class="grid grid-cols-3 bg-spaceGrey h-full">
	<div class="flex flex-col gap-4 px-4 py-8 border-r border-brightGrey">
		<div class="flex justify-between items-center mb-8">
			<h1 class="text-white text-3xl text-ownersTrialWide">Contracts</h1>
			<img src="{PlusWhiteIcon}" alt="Edit icon" class="h-8" />
		</div>

		<div class="h-full overflow-y-scroll">
			{#each Array(5).fill(0) as item}
				<ContractCard
					contactName="Vispute"
					duration="1 week"
					date="21st Feb 2024"
					status="signed"
				/>

				<ContractCard
					contactName="Anubhav"
					duration="2 months"
					date="2nd Jan 2024"
					status="cancelled"
				/>
			{/each}
		</div>
	</div>

	<div class="col-span-2 h-full">
		<div
			class="w-full items-center gap-4 flex p-4 text-ownersTrialWide border-b border-brightGrey"
		>
			<p class="text-white text-xl">New Contract</p>
		</div>
		<div class="py-4 px-8 h-full">
			<h1 class="text-white text-xl signed-badge">
				Transfer of copyrights to label
			</h1>

			<div class="flex flex-row my-8">
				<div class="flex gap-4 has-divider">
					<p class="text-darkWhite">Host:</p>

					<span
						class="rounded-full bg-brightGrey text-darkWhite px-2"
					>
						Aditya
					</span>
				</div>

				<div class="flex gap-4">
					<p class="text-darkWhite">Members:</p>
					<span
						class="rounded-full bg-brightGrey text-darkWhite px-2 mr-[-6px]"
					>
						Aditya
					</span>
					<span
						class="rounded-full bg-brightGrey text-darkWhite px-2 mr-[-6px]"
					>
						Vispute
					</span>
				</div>
			</div>

			<p class="text-white text-sm">
				A music copyright purchase is a legally binding document that
				outlines the terms and conditions of a transaction between a
				buyer and a seller for the sale of a music track. This agreement
				provides clarity and protection to both parties involved inn the
				transaction. Below is a general description of the key
				components typically found in a music record purchase
			</p>

			<div class="flex flex-row my-8">
				<div class="flex gap-4 has-divider">
					<p class="text-darkWhite">Track Name:</p>

					<span
						class="rounded-full border border-brightGrey text-darkWhite px-2 mr-[-6px]"
					>
						The lesser I know, the better
					</span>
				</div>

				<div class="flex gap-4">
					<p class="text-darkWhite">Artist:</p>
					<span
						class="rounded-full bg-brightGrey text-darkWhite px-2 mr-[-6px]"
					>
						Aditya
					</span>
				</div>
			</div>
		</div>
	</div>
</div>

<style lang="postcss">
	.signed-badge::after {
		content: 'signed';
		border: 1px #3bb54a solid;
		border-radius: theme('borderRadius.full');
		font-size: theme('fontSize.sm');
		margin-left: theme('margin.4');
		padding: 0 theme('padding.2');
		color: #3bb54a;
	}

	.has-divider::after {
		content: '';
		border-left: 1px solid theme('colors.brightGrey');
		margin-right: 16px;
	}
</style>
