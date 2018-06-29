const fs = require('fs');

function readSyncByfs(tips) {
    tips = tips || '> ';
    process.stdout.write(tips);
    process.stdin.pause();

    const buf = Buffer.allocUnsafe(10000);
    const response = fs.readSync(process.stdin.fd, buf, 0, 10000, 0);
    process.stdin.resume();
    return buf.toString('utf8',0,response).trim();
}

process.stdout.write('The result is:' + readSyncByfs('请输入任意字符：'));
process.exit(0);
