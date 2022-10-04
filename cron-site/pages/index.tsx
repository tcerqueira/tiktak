import type { NextPage } from 'next'
import Head from 'next/head'
import CronPoster from '../components/index/CronPoster'
import Header from '../components/Header'
import JobList from '../components/index/JobList'

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
