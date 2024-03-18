<script lang="ts">
	import PlusWhiteIcon from '../assets/icons/plus-white.svg';
	import ContractCard from '../lib/components/ContractCard.svelte';
	import { navigate } from 'svelte-routing';
	import { isAuthenticated, signOut } from '../lib/middleware/auth';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!isAuthenticated()) {
			navigate('/');
		}
	});

	let seller = {
		name: '',
		address: '',
		phoneNumber: '',
		email: '',
		governmentID: '',
	};

	let buyer = {
		name: '',
		address: '',
		phoneNumber: '',
		email: '',
		governmentID: '',
	};

	let contract = {
		leaseType: '',
		duration: '',
		validFrom: '',
	};

	function handleSubmit() {
		// Handle form submission
		console.log('Seller Details:', seller);
		console.log('Buyer Details:', buyer);
		console.log('Contract Type:', contract);
	}
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

			<div class="flex gap-4 bg-spaceGrey text-[#A0A0A0] p-4">
				<form class="mx-auto" on:submit|preventDefault="{handleSubmit}">
					<h2 class="text-4xl text-[#785BEB] font-semibold">
						Seller Details
					</h2>
					<div class="grid grid-cols-2 gap-4">
						<div class="grid grid-cols-2 gap-4 py-4">
							<label class="flex flex-col border-brightGrey">
								<span>Name:</span>
								<input
									required
									type="text"
									name="name"
									id="name"
									placeholder="Sample Name"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={seller.name}"
								/>
							</label>
							<label class="flex flex-col">
								<span>Address:</span>
								<input
									required
									type="text"
									name="Address"
									id="Address"
									placeholder="Enter Full Address"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={seller.address}"
								/>
							</label>
							<label class="flex flex-col">
								<span>Phone Number:</span>
								<input
									required
									type="tel"
									name="PhoneNumber"
									id="PhoneNumber"
									placeholder="Enter Phone Number"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={seller.phoneNumber}"
								/>
							</label>
							<label class="flex flex-col">
								<span>Email:</span>
								<input
									required
									type="Email"
									name="Email"
									id="Email"
									placeholder="Enter Email Address"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={seller.phoneNumber}"
								/>
							</label>
						</div>
						<label class="flex flex-col">
							<span class="mb-4">Government ID:</span>
							<div
								class="flex items-center justify-center w-full"
							>
								<label
									for="dropzone-file"
									class="flex flex-col items-center justify-center w-full h-32 border-2 border-[#2B2B2F] border-solid rounded-lg cursor-pointer bg-transparent dark:hover:bg-bray-800 dark:bg-transparent hover:bg-gray-100 dark:transparent dark:hover:border-gray-500 dark:hover:bg-gray-600"
								>
									<div
										class="flex flex-col items-center justify-center pt-5 pb-6"
									>
										<svg
											class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400"
											aria-hidden="true"
											xmlns="http://www.w3.org/2000/svg"
											fill="none"
											viewBox="0 0 20 16"
										>
											<path
												stroke="currentColor"
												stroke-linecap="round"
												stroke-linejoin="round"
												stroke-width="2"
												d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"
											></path>
										</svg>
										<p
											class="mb-2 text-sm text-gray-500 dark:text-gray-400"
										>
											<span class="font-semibold"
												>Click to upload</span
											> or drag and drop
										</p>
										<p
											class="text-xs text-gray-500 dark:text-gray-400"
										>
											SVG, PNG, JPG or GIF (MAX.
											800x400px)
										</p>
									</div>
									<input
										id="dropzone-file"
										type="file"
										class="hidden"
									/>
								</label>
							</div>
						</label>
					</div>

					<h2 class="text-4xl text-[#785BEB] font-semibold">
						Buyer Details
					</h2>
					<div class="grid grid-cols-2 gap-4">
						<div class="grid grid-cols-2 gap-4 py-4">
							<label class="flex flex-col border-brightGrey">
								<span>Name:</span>
								<input
									required
									type="text"
									name="name"
									id="name"
									placeholder="Sample Name"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={buyer.name}"
								/>
							</label>
							<label class="flex flex-col">
								<span>Address:</span>
								<input
									required
									type="text"
									name="Address"
									id="Address"
									placeholder="Enter Full Address"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={buyer.address}"
								/>
							</label>
							<label class="flex flex-col">
								<span>Phone Number:</span>
								<input
									required
									type="tel"
									name="PhoneNumber"
									id="PhoneNumber"
									placeholder="Enter Phone Number"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={buyer.phoneNumber}"
								/>
							</label>
							<label class="flex flex-col">
								<span>Email:</span>
								<input
									required
									type="Email"
									name="Email"
									id="Email"
									placeholder="Enter Email Address"
									class="border border-brightGrey rounded bg-transparent p-2 text-white bind:value={buyer.phoneNumber}"
								/>
							</label>
						</div>

						<label class="flex flex-col">
							<span class="mb-4">Government ID:</span>
							<div
								class="flex items-center justify-center w-full"
							>
								<label
									for="dropzone-file"
									class="flex flex-col items-center justify-center w-full h-32 border-2 border-[#2B2B2F] border-solid rounded-lg cursor-pointer bg-transparent dark:hover:bg-bray-800 dark:bg-transparent hover:bg-gray-100 dark:transparent dark:hover:border-gray-500 dark:hover:bg-gray-600"
								>
									<div
										class="flex flex-col items-center justify-center pt-5 pb-6"
									>
										<svg
											class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400"
											aria-hidden="true"
											xmlns="http://www.w3.org/2000/svg"
											fill="none"
											viewBox="0 0 20 16"
										>
											<path
												stroke="currentColor"
												stroke-linecap="round"
												stroke-linejoin="round"
												stroke-width="2"
												d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"
											></path>
										</svg>
										<p
											class="mb-2 text-sm text-gray-500 dark:text-gray-400"
										>
											<span class="font-semibold"
												>Click to upload</span
											> or drag and drop
										</p>
										<p
											class="text-xs text-gray-500 dark:text-gray-400"
										>
											SVG, PNG, JPG or GIF (MAX.
											800x400px)
										</p>
									</div>
									<input
										id="dropzone-file"
										type="file"
										class="hidden"
									/>
								</label>
							</div>
						</label>
					</div>

					<h2 class="text-4xl text-[#785BEB] font-semibold">
						Contract Type
					</h2>
					<div class="grid grid-cols-2 gap-4 py-4">
						<label class="flex flex-col">
							<span>Lease Type:</span>
							<select
								class="border border-brightGrey rounded p-2 text-white bg-transparent"
								bind:value="{contract.leaseType}"
							>
								<option value="Short">Short Term</option>
								<option value="Long">Long</option>
							</select>
						</label>

						<label class="flex flex-col">
							<span>Duration:</span>
							<input
								class="border border-brightGrey rounded bg-transparent p-2 text-white"
								type="date"
								bind:value="{contract.duration}"
							/>
						</label>
						<label class="flex flex-col">
							<span>Valid From:</span>
							<input
								class="border border-brightGrey rounded bg-transparent p-2 text-white"
								type="date"
								bind:value="{contract.validFrom}"
							/>
						</label>
					</div>
					<div class="flex items-center pt-4">
						<button
							type="submit"
							class="bg-partyPurple text-spaceGrey rounded p-2 mx-auto font-bold uppercase"
						>
							Generate Contract
						</button>
					</div>
				</form>
			</div>
		</div>
	</div>
</div>

<style lang="postcss">
	input[type='date']::-webkit-calendar-picker-indicator {
		filter: invert(100%);
	}
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
