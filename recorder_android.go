//go:build android

package audio

import (
	"gioui.org/app"
	"git.wow.st/gmp/jni"
	"github.com/mearaj/protonet/alog"
)

func checkAndSetRecorderPermission() (err error) {
	defer func() {
		if r := recover(); r != nil {
			alog.Logger().Errorln(r)
		}
		if err != nil {
			alog.Logger().Errorln(err)
		}
	}()
	jvm := jni.JVMFor(app.JavaVM())
	appCtx := jni.Object(app.AppContext())
	err = jni.Do(jvm, func(env jni.Env) (err error) {
		if r := recover(); r != nil {
			alog.Logger().Errorln(r)
		}
		if err != nil {
			alog.Logger().Errorln(err)
		}
		loader := jni.ClassLoaderFor(env, appCtx)
		contextCompatClass, err := jni.LoadClass(env, loader, "androidx/core/content/ContextCompat")
		if err != nil {
			return
		}
		appCtxClass := jni.GetObjectClass(env, appCtx)
		activityMethodID := jni.GetMethodID(env, appCtxClass, "getApplicationContext", "()Landroid/content/Context;")
		activityContext, err := jni.CallObjectMethod(env, appCtx, activityMethodID)
		if err != nil {
			return
		}
		permString := jni.JavaString(env, "android.permission.RECORD_AUDIO")
		reqPermMethodID := jni.GetStaticMethodID(env, contextCompatClass, "checkSelfPermission", "(Landroid/content/Context;Ljava/lang/String;)I")
		//
		res, err := jni.CallStaticIntMethod(env, contextCompatClass, reqPermMethodID, jni.Value(activityContext), jni.Value(permString), jni.Value(1))
		if err != nil {
			alog.Logger().Errorln(err)
		}
		if res == 0 {
			microphonePermissionGranted = true
		} else {
			microphonePermissionGranted = false
		}
		return
	})
	return err
}

// requestPermission calls android's native static method -> androidx.core.app.ActivityCompat.requestPermissions
func requestPermission(view uintptr) (err error) {
	defer func() {
		if err != nil {
			alog.Logger().Errorln(err)
		}
	}()
	jvm := jni.JVMFor(app.JavaVM())
	appCtx := jni.Object(app.AppContext())
	err = jni.Do(jvm, func(env jni.Env) (err error) {
		defer func() {
			if err != nil {
				alog.Logger().Errorln(err)
			}
		}()
		loader := jni.ClassLoaderFor(env, appCtx)
		activityCompatClass, err := jni.LoadClass(env, loader, "androidx/core/app/ActivityCompat")
		if err != nil {
			return
		}
		stringClass, err := jni.LoadClass(env, loader, "java/lang/String")
		if err != nil {
			return
		}
		// org/gioui/GioView can also be used
		gioViewClass, err := jni.LoadClass(env, loader, "android/view/View")
		if err != nil {
			return
		}
		activityMethodID := jni.GetMethodID(env, gioViewClass, "getContext", "()Landroid/content/Context;")
		activityContext, err := jni.CallObjectMethod(env, jni.Object(view), activityMethodID)
		if err != nil {
			return
		}
		permString := jni.JavaString(env, "android.permission.RECORD_AUDIO")
		objArr := jni.NewObjectArray(env, 1, stringClass, jni.Object(permString))
		reqPermMethodID := jni.GetStaticMethodID(env, activityCompatClass, "requestPermissions", "(Landroid/app/Activity;[Ljava/lang/String;I)V")
		err = jni.CallStaticVoidMethod(env, activityCompatClass, reqPermMethodID, jni.Value(activityContext), jni.Value(objArr), jni.Value(1))
		return
	})
	return err
}
