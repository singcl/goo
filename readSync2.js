function read(prompt, callback) {
    process.stdout.write(prompt + ':');
    process.stdin.resume();
    process.stdin.setEncoding('utf-8');
    process.stdin.once('data', function(chunk) {
        process.stdin.pause();
        callback(chunk.trim());
    });
}

read('>>', (chunk) => {
    console.log(chunk);
});

// resume 和'data'监听都会使流进入flowing 模式
