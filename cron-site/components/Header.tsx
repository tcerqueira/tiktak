import Image from 'next/image'
import React from 'react'

function Header() {
	return (
		<header className='fixed w-screen top-0 left-0 flex items-center justify-between mx-auto p-2 bg-orange-300'>
			<div className='relative h-10 w-10'>
				<Image
					src='https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQxeO1p_aGwX1fI7x_iT5hwZUthE60yqAkiNQ&usqp=CAU'
					layout='fill'
				/>
			</div>
			<p>CRON</p>
			<a href="#">About</a>
		</header>
	)
}

export default Header