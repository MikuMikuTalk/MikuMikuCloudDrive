self.importScripts('spark-md5.min.js');

self.onmessage = function(event) {
    const file = event.data;
    const chunkSize = 5 * 1024 * 1024; // 1MB
    const chunks = Math.ceil(file.size / chunkSize);
    const spark = new self.SparkMD5.ArrayBuffer();
    const reader = new FileReader();

    let currentChunk = 0;

    reader.onload = function(e) {
        spark.append(e.target.result);
        currentChunk++;

        if (currentChunk < chunks) {
            loadNextChunk();
        } else {
            const md5 = spark.end();
            self.postMessage({ type: 'complete', md5 });
            self.close();
        }
    };

    reader.onerror = function(e) {
        self.postMessage({ type: 'error', error: 'File read error' });
        self.close();
    };

    function loadNextChunk() {
        const start = currentChunk * chunkSize;
        const end = Math.min(start + chunkSize, file.size);
        const slice = file.slice(start, end);
        reader.readAsArrayBuffer(slice);
    }

    // 发送开始计算的消息
    self.postMessage({ type: 'start' });
    loadNextChunk();
};