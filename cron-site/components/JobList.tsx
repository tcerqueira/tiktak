import React from 'react'
import JobItem from './JobItem'

function JobList() {
  return (
    <div>
        <h1>List</h1>
        <ul>
            <JobItem />
            <JobItem />
            <JobItem />
        </ul>
    </div>
  )
}

export default JobList