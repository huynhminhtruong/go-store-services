Khi bạn thay đổi tên một thư mục trong dự án Go, bạn cần cập nhật lại các import path để chúng phản ánh đúng cấu trúc thư mục mới. Dưới đây là các bước thực hiện:

### Bước 1: Thay Đổi Tên Thư Mục

Đổi tên thư mục theo cấu trúc mong muốn.

### Bước 2: Cập Nhật Các `import` Path

Dò tìm tất cả các file Go có import path chứa tên thư mục cũ và cập nhật lại cho đúng với tên thư mục mới. Bạn có thể dùng lệnh `grep` để tìm kiếm và thay thế nhanh:

```bash
grep -rl 'old_folder_name' . | xargs sed -i 's|old_folder_name|new_folder_name|g'
```

Thay `old_folder_name` và `new_folder_name` bằng tên thư mục cũ và mới.

### Bước 3: Cập Nhật `go.mod`

Chạy lệnh `go mod tidy` để cập nhật các thay đổi trong `go.mod` và loại bỏ các dependencies không còn dùng:

```bash
go mod tidy
```

### Bước 4: Kiểm Tra và Xác Nhận

Kiểm tra lại dự án của bạn bằng cách chạy các lệnh `go build`, `go run`, hoặc `go test` để đảm bảo mọi thứ hoạt động bình thường sau khi cập nhật.
