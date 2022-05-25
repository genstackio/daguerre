package services

import (
	"github.com/genstackio/daguerre/assets"
	"github.com/genstackio/daguerre/commons"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Generate(order *commons.Order) error {

	path, err := generateDotFile(order)

	if nil != err {
		return err
	}

	data, err := generateOutputFromDotFile(path, order)

	if nil != err {
		return err
	}

	_, err = os.Stdout.Write(data)

	return err
}

func generateDotFile(order *commons.Order) (string, error) {
	filename, diag, ctx, err := Build(order)

	if nil != err {
		return "", err
	}

	dir := "./go-diagrams"
	path := dir + "/" + filename + ".dot"

	if err := os.MkdirAll(dir+"/assets/custom", os.ModePerm); err != nil {
		return "", err
	}

	if nil != ctx.CustomAssets {
		for k := range ctx.CustomAssets {
			content, err := assets.Images.ReadFile("images/" + k + ".png")
			if nil != err {
				return "", err
			}
			err = os.WriteFile(dir+"/assets/custom/"+k+".png", content, os.ModePerm)
			if nil != err {
				return "", err
			}

		}
	}
	err = diag.Render()

	if nil != err {
		return "", err
	}

	return path, nil
}

func generateOutputFromDotFile(path string, order *commons.Order) ([]byte, error) {
	format := "png"
	if len(order.Format) > 0 {
		format = order.Format
	}
	cwd, err := os.Getwd()
	if nil != err {
		return []byte{}, err
	}
	err = os.Chdir(filepath.Dir(path))
	if nil != err {
		return []byte{}, err
	}
	c1 := exec.Command("cat", filepath.Base(path))
	c2 := exec.Command("dot", "-T"+format)

	pr, pw := io.Pipe()
	c1.Stdout = pw
	c2.Stdin = pr
	stdout, err := c2.StdoutPipe()

	if nil != err {
		return []byte{}, err
	}
	err1 := c1.Start()

	if nil != err1 {
		defer pw.Close()
		err12 := os.Chdir(cwd)
		if nil != err12 {
			return []byte{}, err12
		}
		return []byte{}, err1
	}
	err2 := c2.Start()

	if nil != err2 {
		defer pw.Close()
		err12 := os.Chdir(cwd)
		if nil != err12 {
			return []byte{}, err12
		}
		return []byte{}, err2
	}

	go func() {
		defer pw.Close()
		err := c1.Wait()
		if nil != err {
			err12 := os.Chdir(cwd)
			if nil != err12 {
				log.Println("c1 command error: " + err12.Error())
				return
			}
			log.Println("c1 command error: " + err.Error())
		}
		err4 := pw.Close()
		if nil != err4 {
			log.Println("pw close error: " + err4.Error())
			return
		}
	}()
	data, _ := io.ReadAll(stdout)
	err = c2.Wait()
	err12 := os.Chdir(cwd)
	if nil != err12 {
		return []byte{}, err12
	}
	if nil != err {
		return []byte{}, err
	}
	return data, nil
}
