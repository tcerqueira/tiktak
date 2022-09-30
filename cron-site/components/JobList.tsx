import React from 'react'
import { CronJob } from '../types/Types'
import JobItem from './JobItem'

const mock: CronJob = {
  id: '12311312321321',
  webhook_url: 'https://localhost:3000/api/',
  webhook_method: 'POST',
  body: 'ITS TIME MF',
  expression: '* * * * *',
  timezone: 'Europe/Lisbon'
};

const mockList: CronJob[] = [
  {...mock}, {...mock}, {...mock}
]

function JobList() {
  return (
    <div>
        <h1>List</h1>
        <ul>
          {mockList.map((item) => {
            return (
              <li key={item.id}>
                <JobItem cronJob={item} />
              </li>
            )
          })}
        </ul>
    </div>
  )
}

export default JobList