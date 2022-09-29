import React, { useState } from 'react'
import TimezoneSelect from 'react-timezone-select'

function CronPoster() {
	const [selectedTimezone, setSelectedTimezone] = useState<any>(Intl.DateTimeFormat().resolvedOptions().timeZone);

	return (
		<div className='p-2 rounded-md border-2 border-orange-500'>
			<h1>Create a CRON job!</h1>
			<form>
				<div className='input-container'>
					<label htmlFor='webhook-in'>Webhook URL</label>
					<input id='webhook-in' type='text' placeholder='https://webhook-example.com/endpoint'/>
				</div>
				<div className='input-container'>
					<label htmlFor='method-in'>Webhook Method</label>
					<input id='method-in' type='text' placeholder='eg: GET, POST, PUT...'/>
				</div>
				<div className='input-container'>
					<label htmlFor='body-in'>Body</label>
					<textarea id="body-in" cols={30} rows={4}>Your time is up!</textarea>
				</div>
				<div className='input-container'>
					<label htmlFor='schedule-in'>Schedule</label>
					<input id='schedule-in' type='text' placeholder='* * * * *'/>
				</div>
				<div className='input-container'>
					<label htmlFor='timezone-in'>Timezone</label>
					<TimezoneSelect value={selectedTimezone} onChange={setSelectedTimezone} />
				</div>
			</form>
		</div>
	)
}

export default CronPoster