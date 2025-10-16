import axios from 'axios';

const API_BASE_URL = 'http://localhost:9090'; // Replace with your API base URL

const axiosPrivate = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',

  },
  withCredentials: true, // important for HTTP-only cookies
});

export default axiosPrivate;