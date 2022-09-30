import React, { useState } from 'react'
import TimezoneSelect from 'react-timezone-select'
import { ClockIcon } from '@heroicons/react/solid'

function CronPoster() {
	const [selectedTimezone, setSelectedTimezone] = useState<any>(Intl.DateTimeFormat().resolvedOptions().timeZone);

	return (
		<div className='p-2 rounded-md border-2 border-orange-500'>
			<h1>Create a CRON job!</h1>
			<form className='flex flex-col'>
				<div className='md:flex md:items-center'>
					<div className='input-container md:grow'>
						<label className='md:mr-8' htmlFor='webhook-in'>Webhook URL</label>
						<input id='webhook-in' type='text' placeholder='https://webhook-example.com/endpoint'/>
					</div>
					<div className='input-container'>
						<label htmlFor='method-in'>Webhook Method</label>
						<input className='md:w-[1rem]' id='method-in' type='text' placeholder='eg: GET, POST, PUT...'/>
					</div>
				</div>
				<div className='input-container'>
					<label htmlFor='body-in'>Body</label>
					<textarea id="body-in" cols={30} rows={4} defaultValue='Your time is up!' />
				</div>
				<div className='input-container'>
					<label htmlFor='schedule-in'>Schedule</label>
					<input id='schedule-in' type='text' placeholder='* * * * *'/>
				</div>
				<div className='input-container'>
					<label htmlFor='timezone-in'>Timezone</label>
					<TimezoneSelect value={selectedTimezone} onChange={selectedTimezone} />
				</div>
				<div className='flex justify-center items-center md:w-[60%] w-[100%] rounded-lg self-center cursor-pointer hover:backdrop-brightness-90'
				onClick={() => console.log('yo')}>
					<ClockIcon className='h-10 w-10 text-orange-600' />
					<button type='submit' hidden />
				</div>
			</form>
		</div>
	)
}

export default CronPoster