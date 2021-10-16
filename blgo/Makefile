default:
	tmux_send "make build"

build: clean
	@mkdir frames
	@go run main.go

clean:
	@rm -rf frames out.png out.gif out.mp4 palette.png

video:
	ffmpeg -framerate 30 -i frames/frame_%04d.png -s 1920x1080 -c:v libx264 -preset veryslow -crf 20 -pix_fmt yuv420p -vf "fade=t=out:st=55:d=5" out.mp4
	# ffmpeg -framerate 30 -i frames/frame_%04d.png -s 1920x1080 -c:v libx264 -preset veryslow -crf 20 -pix_fmt yuv420p out.mp4
	# ffmpeg -framerate 30 -i frames/frame_%04d.png -s 1280x720 -c:v libx264 -preset veryslow -crf 20 -pix_fmt yuv420p -vf "fade=t=out:st=55:d=5" out.mp4
	# ffmpeg -framerate 30 -i frames/frame_%04d.png -s 1280x720 -c:v libx264 -preset veryslow -crf 20 -pix_fmt yuv420p out.mp4

audio:
	ffmpeg -i out.mp4 -i audio.wav -map 0:v -map 1:a -c:v copy final.mp4
	# if audio is longer than video:
	# ffmpeg -i out.mp4 -i audio.wav -map 0:v -map 1:a -c:v copy -shortest final.mp4

