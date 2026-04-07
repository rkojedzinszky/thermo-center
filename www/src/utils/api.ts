import fetch from './fetch'
import { DefaultApi, Configuration } from '@/api'

const configuration = new Configuration({ basePath: window.location.origin, fetchApi: fetch })
export default new DefaultApi(configuration)
