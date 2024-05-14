import axios, { AxiosResponse } from 'axios';

interface UploadResponse {
  success: boolean;
  message: string;
  urls?: string[];
}

async function uploadFiles(files: FileList): Promise<UploadResponse> {
  try {
    const formData = new FormData();
    Array.from(files).forEach((file, index) => {
      formData.append(`file${index}`, file);
    });

    const response: AxiosResponse<UploadResponse> = await axios.post('/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    return response.data;
  } catch (error) {
    return {
      success: false,
      message: 'An error occurred while uploading the files.',
    };
  }
}

export default uploadFiles;
