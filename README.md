# golang RSA加解密实例
练手项目

## RSA公私钥生成过程
1. 生成私钥
    ```bash
    openssl genrsa -out privacy.key 2048
    ```
2. 生成公钥
    ```bash
        openssl -in privacy.key -pubout -out public.key
    ```
