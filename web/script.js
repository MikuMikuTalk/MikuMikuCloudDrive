const fileInput = document.getElementById('file-input');
const uploadBtn = document.getElementById('upload-btn');
const progressBar = document.getElementById('progress-bar');
const progress = document.getElementById('progress');
const statusText = document.getElementById('status');
const token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjMsInVzZXJuYW1lIjoidmVkYWw5ODciLCJlbWFpbCI6InZlZGFsOTg3QGV4YW1wbGUuY29tIiwiZXhwIjoxNzM3NDIyMjMwLCJqdGkiOiIwY2Q3YWNiYy04YTIxLTQwMjgtYmViYi1hNjRkN2Q4NTI5MjIifQ.ohGaCOcxytnvM549YMJ_ZFTbMUeFYhl1wWa7h0Ll4zY"
let file;
fileInput.addEventListener('change', (event) => {
    file = event.target.files[0];
    if (file) {
        statusText.textContent = `选择文件:${file.name}`;
    } else {
        statusText.textContent = '未选择文件';
    }

});
uploadBtn.addEventListener('click', async () => {
    try {
        if (!file) {
            alert('请选择文件');
            return;
        }
        // 显示MD5计算提示
        document.getElementById('md5-progress').textContent = '正在计算文件MD5，请稍候...';
        // 使用setTimeout让浏览器有时间渲染DOM
        await new Promise(resolve => setTimeout(resolve, 0));

        const chunkSize = 5 * 1024 * 1024; 
        const chunks = sliceFile(file, chunkSize);
        let md5 = await calculateFileMD5(file);

        // 隐藏MD5计算提示
        document.getElementById('md5-progress').textContent = '';

        const uploadedChunks = await getUploadedChunks(file, md5, chunks.length, 'http://127.0.0.1:8888/file/getUploadedChunks');
        console.log(uploadedChunks);
        
        await uploadChunks(file, chunks, md5, 'http://127.0.0.1:8888/file/upload', (uploaded, total) => {
            const percent = (uploaded / total) * 100;
            progress.style.width = `${percent}%`;
            statusText.textContent = `上传中: ${uploaded}/${total} 分片 (${percent.toFixed(2)}%)`;
        }, uploadedChunks);
        await mergeChunks(file, chunks.length, md5, 'http://127.0.0.1:8888/file/merge');
        statusText.textContent = '上传完成!';
        progress.style.width = '100%';
    } catch (error) {
        console.error('上传过程中发生错误:', error);
        statusText.textContent = '上传失败，请重试';
        // 隐藏MD5计算提示
        document.getElementById('md5-progress').textContent = '';
    }
});

/**
 * Slices a file into smaller chunks of a specified size.
 *
 * @param {File} file - The file to be sliced.
 * @param {number} chunkSize - The size of each chunk in bytes.
 * @returns {Array<Blob>} An array of Blob objects representing the file chunks.
 *
 */
function sliceFile(file, chunkSize) {
    const chunks = []
    let start = 0;
    while (start < file.size) {
        const end = start + chunkSize;
        const chunk = file.slice(start, end);
        chunks.push(chunk);
        start = end;
    }
    return chunks;
}

async function uploadChunks(file, chunks, md5, url, onProgress, uploadedChunks) {
    const totalChunks = chunks.length;
    let uploaded = 0;

    for (let i = 0; i < totalChunks; i++) {
        if (uploadedChunks.includes(i)) {
            onProgress(++uploaded, totalChunks);
            continue;
        }

        const chunk = chunks[i];
        const formData = new FormData();
        formData.append('file', chunk, file.name);
        formData.append('chunkIndex', i);
        formData.append('totalChunks', totalChunks);
        formData.append('fileMD5', md5);

        let retryCount = 10; // 重试次数
        while (retryCount > 0) {
            try {
                const response = await fetch(url, {
                    headers:{
                        'Authorization':token
                    },
                    method: 'POST',
                    body: formData,
                });
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                if (data.code === 200) {
                    onProgress(++uploaded, totalChunks);
                    break; // 上传成功，跳出重试循环
                } else {
                    console.error('上传分片失败:', data.message);
                    retryCount--;
                }
            } catch (error) {
                console.error('上传分片时发生错误:', error);
                retryCount--;
                if(retryCount > 0) {
                    await new Promise(resolve => setTimeout(resolve, 1000)); // 可选：添加延迟以避免快速连续请求
                }
            }
            if (retryCount === 0) {
                throw error; // 重试次数用尽，抛出错误
            }
        }
    }
}
function calculateFileMD5(file) {
    return new Promise((resolve, reject) => {
        const worker = new Worker('md5-worker.js');
        worker.onmessage = (event) => {
            if (event.data.type === 'complete') {
                resolve(event.data.md5);
            } else if (event.data.type === 'error') {
                reject(event.data.error);
            }
        };
        worker.onerror = (error) => {
            reject(error);
        };
        worker.postMessage(file);
    });
}


async function mergeChunks(file, totalChunks, md5, url) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization':token
        },
        body: JSON.stringify({
            directory_id: 5,
            filename: file.name,
            totalChunks: totalChunks,
            fileMD5: md5
        })
    })
    const data = await response.json();
    if (!data.code === 200) {
        console.error('文件合并失败:', data.message);
    }
}


async function getUploadedChunks(file, md5, totalChunks, url) {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization':token
            },
            body: JSON.stringify({
                filename: file.name,
                fileMd5: md5,
                totalChunks: totalChunks,
            }),
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        if (data.code === 200) {

            return data.data.chunksArray;
        } else {
            console.error('获取已上传分片失败:', data.message);
            return [];
        }
    } catch (error) {
        console.error('获取已上传分片时发生错误:', error);
    }
}