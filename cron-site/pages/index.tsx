import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import CronPoster from '../components/CronPoster'
import Header from '../components/Header'
import JobList from '../components/JobList'

const Home: NextPage = () => {
  return (
    <div className='bg-orange-200'>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />
      <main className='min-h-screen pt-16 px-1 max-w-[768px] mx-auto'>
        <CronPoster />
        <JobList />
      </main>
    </div>
  )
}

export default Home
