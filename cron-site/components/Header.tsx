import Image from 'next/image'
import React from 'react'
import { CalendarIcon } from '@heroicons/react/solid'

function Header() {
	return (
		<header className='fixed w-screen top-0 left-0 bg-orange-300'>
			<div className='flex items-center justify-between mx-auto p-2 md:w-[768px]'>
				<div className='relative h-10 w-20'>
					{/* <Image
						src='https://w7.pngwing.com/pngs/800/101/png-transparent-vixie-cron-linux-command-execution-linux-text-logo-command-thumbnail.png'
						layout='fill'
					/> */}
					<CalendarIcon className='h-10 w-10 text-orange-800'/>
				</div>
				<p className='font-bold text-2xl cursor-default'>
					<span className='text-green-600'>C</span>
					<span className='text-blue-600'>R</span>
					<span className='text-purple-600'>O</span>
					<span className='text-red-600'>N</span>
				</p>
				<a className='hover:underline hover:underline-offset-4' href="https://github.com/tcerqueira/tiktak" target='_blank'>Code</a>
			</div>
		</header>
	)
}

export default Header