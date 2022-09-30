import React from 'react'
import { CronJob } from '../types/Types'
import JobItem from './JobItem'

const mock: CronJob = {
  id: '12311312321321',
  webhook_url: 'https://localhost:3000/api/',
  webhook_method: 'POST',
  body: 'ITS TIME MF',
  expression: '* * * * *',
  timezone: 'WEST'
};

function JobList() {
  return (
    <div>
        <h1>List</h1>
        <ul>
            <li>
              <JobItem cronJob={mock} />
            </li>
            <li>
              <JobItem cronJob={mock} />
            </li>
            <li>
              <JobItem cronJob={mock} />
            </li>
        </ul>
    </div>
  )
}

export default JobList