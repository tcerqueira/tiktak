import React, { useCallback, useState } from 'react'
import TimezoneSelect, { ITimezone } from 'react-timezone-select'
import { ClockIcon } from '@heroicons/react/solid'
import { useForm } from 'react-hook-form';
import TextInput from '../TextInput';
import TextArea from '../TextArea';
import SelectInput from '../SelectInput';

interface FormData {
	webhook_url: string;
	webhook_method: string;
	body: string;
	expression: string;
}

function CronPoster() {
	const [selectedTimezone, setSelectedTimezone] = useState<ITimezone>(Intl.DateTimeFormat().resolvedOptions().timeZone);
	const { register, handleSubmit } = useForm<FormData>()

	const onSubmit = useCallback((data: FormData) => {
		console.log(
			{ ...data, selectedTimezone}
		)
	}, [selectedTimezone])

	return (
		<div className='p-2 rounded-md border-2 border-orange-500'>
			<h1>Create a CRON job!</h1>
			<form className='flex flex-col' onSubmit={handleSubmit(onSubmit)}>
				<div className='webhook-container'>
					<div className='input-container md:grow'>
						<TextInput id='webhook-in' label='Webhook URL' placeholder='https://webhook-example.com/endpoint' {...register('webhook_url')}/>
					</div>
					<div className='input-container md:basis-1'>
						<SelectInput id='method-in' label='Webhook Method' {...register('webhook_method')}>
							<option value="GET">GET</option>
							<option value="POST">POST</option>
							<option value="PUT">PUT</option>
							<option value="PATCH">PATCH</option>
							<option value="DELETE">DELETE</option>
						</SelectInput>
					</div>
				</div>
				<div className='input-container'>
					<TextArea id='body-in' label='Body' defaultValue='Your time is up!' {...register('body')}/>
				</div>
				<div className='input-container'>
					<TextInput id='schedule-in' label='Schedule' placeholder='* * * * *' {...register('expression')}/>
				</div>
				<div className='input-container'>
					<label htmlFor='timezone-in'>Timezone</label>
					<TimezoneSelect value={selectedTimezone} onChange={(tz) => {
							setSelectedTimezone(tz);
							console.log(tz);
						}}
					/>
				</div>
				<div className='flex justify-center items-center md:w-[60%] w-[100%] rounded-lg self-center cursor-pointer hover:backdrop-brightness-90'
				onClick={() => console.log('submitted')}>
					<ClockIcon className='h-10 w-10 text-orange-600' />
					<button type='submit' hidden />
				</div>
			</form>
		</div>
	)
}

export default CronPoster