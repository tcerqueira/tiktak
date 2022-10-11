import React, { useCallback, useRef, useState } from 'react'
import TimezoneSelect, { ITimezone } from 'react-timezone-select'
import { ClockIcon } from '@heroicons/react/solid'
import { useForm } from 'react-hook-form';
import TextInput from '../TextInput';
import TextArea from '../TextArea';
import SelectInput from '../SelectInput';
import { cronRegex } from '../../utils/RegEx'
import { postCron, PostCronPayload } from '../../utils/ApiCalls';

export interface CronFormData {
	webhook_url: string;
	webhook_method: string;
	body: string;
	cron_expression: string;
};

interface CronPosterProps {
	onPost?: (data: CronFormData) => void;
}

function CronPoster({ onPost }: CronPosterProps) {
	const [selectedTimezone, setSelectedTimezone] = useState<ITimezone>(Intl.DateTimeFormat().resolvedOptions().timeZone);
	const { register, handleSubmit, formState: {errors} } = useForm<CronFormData>({
		defaultValues: {
			webhook_url: 'http://localhost:8050/webhook',
			webhook_method: 'POST',
			body: "Time's up!",
			cron_expression: '* * * * *',
		}
	});

	const onSubmit = useCallback(async (data: CronFormData) => {
		try {
			const timezone = typeof selectedTimezone === 'string' ? selectedTimezone : selectedTimezone.value;
			const payload: PostCronPayload = { ...data, timezone };
			const response = await postCron(payload);
			onPost && onPost(data);
		} catch (err) {
			console.error(err)
		}

	}, [selectedTimezone])

	return (
		<div className='p-2 rounded-md border-2 border-orange-500'>
			<h1>Create a CRON job!</h1>
			<form className='flex flex-col' onSubmit={handleSubmit(onSubmit)}>
				<div className='webhook-container'>
					<div className='input-container md:grow'>
						<TextInput id='webhook-in' label='Webhook URL' placeholder='https://webhook-example.com/endpoint'
							{...register('webhook_url', { required: true })}/>
					</div>
					<div className='input-container md:basis-1'>
						<SelectInput label='Webhook Method'
							{...register('webhook_method', { required: true })}>
							<option value="POST">POST</option>
							<option value="GET">GET</option>
							<option value="PUT">PUT</option>
							<option value="PATCH">PATCH</option>
							<option value="DELETE">DELETE</option>
						</SelectInput>
					</div>
				</div>
				{ errors.webhook_url && <p className='error-message'>Webhook URL is required.</p> }
				<div className='input-container'>
					<TextArea id='body-in' label='Body' defaultValue='Your time is up!'
						{...register('body')}/>
				</div>
				<div className='input-container'>
					<TextInput id='schedule-in' label='Schedule' placeholder='* * * * *'
						{...register('cron_expression', { required: true, pattern: cronRegex })}/>
				</div>
				{ errors.cron_expression && <p className='error-message'>Invalid CRON expression.</p> }
				<div className='input-container'>
					<label htmlFor='timezone-in'>Timezone</label>
					<TimezoneSelect value={selectedTimezone} onChange={(tz) => {
							setSelectedTimezone(tz);
							console.log(tz);
						}}
					/>
				</div>
				<button type='submit' className='submit-btn bg-teal-200 mt-4'>
					<ClockIcon className='h-10 w-10 text-orange-600' />
				</button>
			</form>
		</div>
	)
}

export default CronPoster