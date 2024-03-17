<script lang="ts">
	import UnmuteLogo from '../assets/logo.svg';
	import BrandGoogle from '../assets/icons/brand-google.svg';
</script>

<div
	class="w-full min-h-screen bg-spaceGrey flex flex-col items-center justify-center gap-8 p-8"
>
	<img
		src="{UnmuteLogo}"
		alt="Unmute Logo"
		class="absolute top-8 md:left-8 right-8"
	/>
	<div class="w-full sm:w-2/3 lg:w-1/3 text-white">
		<a href="/login" class="uppercase text-darkWhite">Login</a><span
			class="border-l mx-3"
		></span><span class="uppercase text-partyPurple">Sign Up</span>
	</div>

	<div class="flex flex-col gap-2 w-full sm:w-2/3 lg:w-1/3">
		<h1 class="text-white text-2xl">Welcome,</h1>
		<p class="text-darkWhite">Please enter your details below to sign up</p>
	</div>

	<form action="" class="flex flex-col w-full sm:w-2/3 lg:w-1/3">
		<label for="name" class="text-white text-sm mb-2 uppercase">
			Full Name
		</label>
		<input
			type="text"
			name="name"
			id="name"
			placeholder="Ramesh Waterwala"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2"
		/>

		<label for="email" class="text-white text-sm mt-8 mb-2 uppercase">
			Email
		</label>
		<input
			type="email"
			name="email"
			id="email"
			placeholder="example@gmail.com"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2"
		/>

		<label for="password" class="text-white text-sm mt-8 mb-2 uppercase">
			Password
		</label>
		<input
			type="password"
			name="password"
			id="password"
			placeholder="******"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2"
		/>

		<label
			for="confpassword"
			class="text-white text-sm mt-8 mb-2 uppercase"
		>
			Confirm Password
		</label>
		<input
			type="password"
			name="confpassword"
			id="confpassword"
			placeholder="******"
			class="border border-brightGrey active:border-partyPurple active:ring-partyPurple rounded bg-transparent p-2 mb-8"
		/>

		<button
			type="submit"
			class="bg-partyPurple text-spaceGrey rounded py-2 font-bold uppercase"
		>
			Sign Up
		</button>
	</form>

	<button
		class="flex justify-center items-center gap-2 border border-partyPurple text-darkWhite bg-transparent hover:bg-partyPurple w-full sm:w-1/2 lg:w-1/3 rounded py-2 font-bold uppercase"
	>
		<img src="{BrandGoogle}" alt="Google Icon" class="h-4" />
		<span> Sign in with Google </span>
	</button>
</div>
