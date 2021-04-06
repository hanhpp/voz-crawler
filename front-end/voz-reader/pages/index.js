import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Footer from '../components/footer'
import Body from '../components/body'

const App = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Body styles={styles}>
          <p>Hello</p>
      </Body>

      <Footer className={styles.footer}></Footer>
    </div>
  )
}

export default App

