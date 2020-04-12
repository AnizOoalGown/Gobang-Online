module.exports = {
    devServer: {
        port: 5555,
        proxy: {
            '/ws': {
                target: 'http://localhost:8080',
                ws: true,
                changeOrigin: true,
            },
        }
    },
}